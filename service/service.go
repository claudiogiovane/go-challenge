package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	listener := flag.String("http.address", ":8080", "http listen address")
	flag.Parse()
	http.HandleFunc("/numbers", func(w http.ResponseWriter, r *http.Request) { handler(w, r) })
	log.Fatal(http.ListenAndServe(*listener, nil))
}

//Function handler to resolve the request from the service
func handler(w http.ResponseWriter, r *http.Request) {
	maxDelay := 500 * time.Millisecond
	ctx, quit := context.WithTimeout(context.Background(), maxDelay)
	defer quit()
	params := r.URL.Query()
	//query parameters must be called after 'u'
	qparamurl := r.URL.Query()["u"]
	// in case the query parameters are not informed, return status 400
	if len(qparamurl) == 0 {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request: no parameters was informed"))
		return
	}
	// in case the query parameters informed is not 'u', return status 400
	keys := []string{}
	for key := range params {
		keys = append(keys, key)
		if len(keys) > 1 {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request: Only 'u' is accepted as a param"))
			return
		}
	}
	// creating the channel for retrieve the numbers independently
	channel := make(chan []int, len(qparamurl))
	// retrieve all URLs  and working on the data concurrently
	for i := range qparamurl {
		go func(returl string) {
			//retrieve the numbers
			numbers, numerr := retriNumbers(ctx, returl)
			if numerr != nil {
				log.Printf("Invalid URL - : %s", numerr)
				return
			} else if numerr == nil && numbers != nil && len(numbers) > 0 {
				//reorganize the numbers and retrieve then sorted and without repetiton
				reorgnum, err := Reorg(ctx, numbers)
				if err != nil {
					log.Printf("Unable to reorganize the numbers - : %s", err)
					return
				}
				channel <- reorgnum //sending the data through the channel
			}
		}(qparamurl[i])
	}
	// the final slice result will be stored here
	var sortedslice = []int{}
	// counter for the merge operations
	i := 0
	// process merge of slices sorted in the URL parameters
Outer:
	for {
		select {
		//case the request take more than 500ms, abort and deliver only set of numbers
		//retrieved in the interval
		case <-ctx.Done():
			log.Printf("Could not process all numbers. Timeout request failed to complete in 500ms")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"numbers": sortedslice})
			return
		//Otherwise, continue to process the merge
		case res := <-channel: //receinving the data 
			sortedslice = Mergenumbers(sortedslice, res)
			i++
			if i >= len(qparamurl) {
				break Outer
			}
		}
	}

	log.Printf("All the numbers were retrieved and organized")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"numbers": sortedslice})
}
