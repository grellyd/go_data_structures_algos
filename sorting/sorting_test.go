package sorting

import (
	"testing"
	"fmt"
)

func TestMergesort(t *testing.T) {
	var tests = []struct {
		input []int
		output []int
	} {
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1,2,3},
			[]int{1,2,3},
		},
		{
			[]int{3,2,1},
			[]int{1,2,3},
		},
		{
			[]int{3,1,2},
			[]int{1,2,3},
		},
		{
			[]int{3,7,1,2,4,9,5,10,8,6},
			[]int{1,2,3,4,5,6,7,8,9,10},
		},
		{ 
			[]int{-9,-4,-2,-7},
			[]int{-9,-7,-4,-2},
		},
	}
	for _, test := range tests {
		result := Mergesort(test.input)
		fmt.Printf("Result: %v\n\n", result)
		for i, term := range result {
			if test.output[i] != term {
				t.Errorf("TestMergesort Error! %v doesn't match %v!", result, test.output)
			}
		}
	}
}

func TestBubblesort(t *testing.T) {
	var tests = []struct {
		input []int
		output []int
	} {
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1,2,3},
			[]int{1,2,3},
		},
		{
			[]int{3,2,1},
			[]int{1,2,3},
		},
		{
			[]int{3,1,2},
			[]int{1,2,3},
		},
		{
			[]int{3,7,1,2,4,9,5,10,8,6},
			[]int{1,2,3,4,5,6,7,8,9,10},
		},
		{ 
			[]int{-9,-4,-2,-7},
			[]int{-9,-7,-4,-2},
		},
	}
	for _, test := range tests {
		result := Bubblesort(test.input)
		fmt.Printf("Result: %v\n\n", result)
		for i, term := range result {
			if test.output[i] != term {
				t.Errorf("TestMergesort Error! %v doesn't match %v!", result, test.output)
			}
		}
	}
}


