package main

import (
	"context"
	"fmt"
	"sort"
)

// Reorg function sorts the slice and eliminate duplicate numbers of it.
func Reorg(ctx context.Context, list []int) (res []int, err error) {
	//sorting
	array, err := sortnumbers(ctx, list)
	if err != nil {
		return nil, err
	}
	//eliminates repeated numbers
	notrepeated := removeDuplicates(array)

	return notrepeated, nil
}

// sortnumbers take the array and returns the sorted numbers in ascending order
func sortnumbers(ctx context.Context, list []int) (res []int, err error) {
	sorted := append([]int{}, list...)

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("URL timeout")
	default:
		sort.Ints(sorted)
	}

	return sorted, nil
}

// The slice returned by removeDuplicates has all duplicates removed,
func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// Mergenumbers concat the request's result and operates the final reorg.
func Mergenumbers(first []int, second []int) []int {
	merged := []int{}

	//concatenate the slices
	merged = append(first, second...)

	//removes duplicate and sort the final slice
	merged = removeDuplicates(merged)
	sort.Ints(merged)

	return merged
}
