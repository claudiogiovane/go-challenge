package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Numbers is a struct for store the parsed JSON returns.
type Numbers struct{ Numbers []int }

//validURL retrieve each URLs if they are syntactically valid
func validURL(geturl string) error {
	u, err := url.Parse(geturl)

	if err != nil {
		return err
	}

	switch {
	case u.Scheme != "https" && u.Scheme != "http":
		return fmt.Errorf("Error: address must begin with http or https in %s", geturl)
	case u.Scheme == "" || u.Host == "":
		return fmt.Errorf("Error: address must be an absolute URL in %s", geturl)
	default:
		return nil
	}
}

// parserJSON is a Parser for the received JSON
func parserJSON(body []byte) (res []int, err error) {
	var num Numbers
	json.Unmarshal(body, &num)
	return num.Numbers, nil
}

// retriNumbers retrieves the numbers from the request
func retriNumbers(ctx context.Context, geturl string) (resp []int, err error) {
	// validate url
	urlerr := validURL(geturl)
	if urlerr != nil {
		return nil, urlerr
	}

	// retrieve response
	client := &http.Client{}
	req, reqerr := http.NewRequest("GET", geturl, nil)
	if reqerr != nil {
		return nil, reqerr
	}

	req = req.WithContext(ctx)
	r, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP: Status code shoud be 200, but is %d",
			r.StatusCode)
	}

	defer r.Body.Close()

	body, bodyerr := ioutil.ReadAll(r.Body)
	if bodyerr != nil {
		return nil, bodyerr
	}
	//get the numbers from the JSON result
	return parserJSON(body)
}
