package sorting

import (
	"testing"
)
	
var tests = []struct {
	input []int
	output []int
} {
	{
		// empty
		[]int{},
		[]int{},
	},
	{
		// single
		[]int{1},
		[]int{1},
	},
	{
		// no change
		[]int{1,2,3},
		[]int{1,2,3},
	},
	{
		// inverted odd
		[]int{3,2,1},
		[]int{1,2,3},
	},
	{
		// inverted even
		[]int{4,3,2,1},
		[]int{1,2,3,4},
	},
	{
		// with zero
		[]int{3,0,1,2},
		[]int{0,1,2,3},
	},
	{
		// with negative
		[]int{3,0,1,-1,2},
		[]int{-1,0,1,2,3},
	},
	{
		// many
		[]int{3,7,1,2,-1,4,9,5,10,8,6},
		[]int{-1,1,2,3,4,5,6,7,8,9,10},
	},
	{ 
		// all negative
		[]int{-9,-4,-2,-7},
		[]int{-9,-7,-4,-2},
	},
}

func TestQuicksort(t *testing.T) {
	for _, test := range tests {
		result := Quicksort(test.input)
		if len(result) != len(test.output) {
			t.Errorf("TestQuicksort Error! len(%v) doesn't match len(%v)!", result, test.output)
		}
		for i, term := range result {
			if test.output[i] != term {
				t.Errorf("TestQuicksort Error! %v doesn't match %v!", result, test.output)
			}
		}
	}
}


func TestMergesort(t *testing.T) {
	for _, test := range tests {
		result := Mergesort(test.input)
		if len(result) != len(test.output) {
			t.Errorf("TestMergesort Error! len(%v) doesn't match len(%v)!", result, test.output)
		}

		for i, term := range result {
			if test.output[i] != term {
				t.Errorf("TestMergesort Error! %v doesn't match %v!", result, test.output)
			}
		}
	}
}


func TestBubblesort(t *testing.T) {
	for _, test := range tests {
		result := Bubblesort(test.input)
		if len(result) != len(test.output) {
			t.Errorf("TestBubblesort Error! len(%v) doesn't match len(%v)!", result, test.output)
		}
		for i, term := range result {
			if test.output[i] != term {
				t.Errorf("TestBubblesort Error! %v doesn't match %v!", result, test.output)
			}
		}
	}
}


