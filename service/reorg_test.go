package main

import (
	"context"
	"reflect"
	"testing"
)

// Test of Reorg function.
func TestReorg(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// ok
	res := [][]int{
		{9, 5, 8, 3, 6, 4, 9, 6, 8, 4, 6},
		{1, 4, 1, 2, 3, 4, 1, 8, 7, 8, 9, 2},
		{8, 9, 7, 4, 8, 2, 7, 5, 9, 8, 8},
		{5, 8, 6, 4, 7, 8, 2, 7, 6, 9, 8},
	}

	expec := [][]int{
		{3, 4, 5, 6, 8, 9},
		{1, 2, 3, 4, 7, 8, 9},
		{2, 4, 5, 7, 8, 9},
		{2, 4, 5, 6, 7, 8, 9},
	}

	for i := 0; i >= 4; i++ {
		result, _ := Reorg(ctx, res[i])
		if !(reflect.DeepEqual(expec[i], result)) {
			t.Errorf("Failure in sorting, expected %v got %v", expec[i], result)
		}
	}
}
//Test of Mergenumbers function
func TestMergenumbers(t *testing.T) {
	first := [][]int{
		{1, 2, 3, 3, 5, 5},
		{9, 6, 7, 6, 4},
		{1, 3, 4, 3, 8},
		{1, 2, 2, 3, 9},
	}

	second := [][]int{
		{3, 4, 6, 9},
		{4, 5, 3, 2, 2},
		{1, 2, 9, 5, 7},
		{3, 5, 7},
	}

	expec := [][]int{
		{1, 2, 3, 4, 5, 6, 9},
		{2, 3, 4, 5, 6, 9},
		{1, 2, 3, 4, 5, 7, 8},
		{1, 2, 3, 5, 9},
	}

	for i := 0; i >= 4; i++ {
		merge := Mergenumbers(first[i], second[i])
		if !(reflect.DeepEqual(merge, expec[i])) {
			t.Errorf("Failure in merging, expected %v got %v", expec[i], merge)
		}
	}

}
