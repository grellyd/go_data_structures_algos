package fib_test

import (
	"testing"
	"fib"
)

func TestComputeNextValue(t *testing.T) {
	var tests = []struct {
		input []int
		output int
	}{
		{fib.FibBase, 1},
		{append(fib.FibBase, 1), 2},
		{[]int{0, 1, 1, 2}, 3},
		{[]int{0, 1, 1, 2, 3, 5, 8}, 13},

	}
	for _, test := range tests {
		if result := fib.ComputeNextValue(test.input); result != test.output {
			t.Errorf("ComputeNextValue(%q) = %v", test.input, result)
		}
	}
}

func TestAddNextValue(t *testing.T) {
	var tests = []struct {
		input []int
		output []int
	}{
		{fib.FibBase, []int{0, 1, 1}},
		{append(fib.FibBase, 1), []int{0, 1, 1, 2}},
		{[]int{0, 1, 1, 2}, []int{0, 1, 1, 2, 3}},
		{[]int{0, 1, 1, 2, 3, 5, 8}, []int{0, 1, 1, 2, 3, 5, 8, 13}},

	}
	for _, test := range tests {
		result := fib.AddNextValue(test.input)
		outcome := true
		if len(test.output) == len(result) {
			for i, val := range result {
				outcome = outcome && val == test.output[i]
			}
		} else {
			outcome = false
		}
		if !outcome {
			t.Errorf("AddNextValue(%v) = %v instead of %v", test.input, result, test.output)
		}
	}
}

func TestGetFibs(t *testing.T) {
	var tests = []struct {
		input int
		output []int
	}{
		{0, fib.FibBase},
		{1, fib.FibBase},
		{2, fib.FibBase},
		{4, []int{0, 1, 1, 2}},
		{5, []int{0, 1, 1, 2, 3}},
		{8, []int{0, 1, 1, 2, 3, 5, 8, 13}},

	}
	for _, test := range tests {
		result := fib.GetFibs(test.input)
		outcome := true
		if len(test.output) == len(result) {
			for i, val := range result {
				outcome = outcome && val == test.output[i]
			}
		} else {
			outcome = false
		}
		if !outcome {
			t.Errorf("GetFibs(%v) = %v instead of %v", test.input, result, test.output)
		}
	}
}

func TestIsFib(t *testing.T) {
	var tests = []struct {
		input int
		output bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{13, true},
		{-1, false},
		{4, false},
		{15, false},
	}
	for _, test := range tests {
		if result := fib.IsFib(test.input); result != test.output {
			t.Errorf("IsFib(%v) = %v instead of %v", test.input, result, test.output)
		}
	}
}

func TestIsFibs(t *testing.T) {
	var tests = []struct {
		input []int
		output bool
	}{
		{[]int{0}, true},
		{[]int{1, 0}, true},
		{[]int{2, 3, 4}, false},
		{[]int{13, 21}, true},
		{[]int{-1, 0, 1, 1}, false},
		{[]int{4, 15}, false},
	}
	for _, test := range tests {
		if result := fib.IsFibs(test.input); result != test.output {
			t.Errorf("IsFibs(%v) = %v instead of %v", test.input, result, test.output)
		}
	}
}

func TestNthFib(t *testing.T) {
	var tests = []struct {
		input int
		output int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{8, 13},
		{-1, 0},
	}
	for _, test := range tests {
		if result := fib.NthFib(test.input); result != test.output {
			t.Errorf("NthFib(%v) = %v instead of %v", test.input, result, test.output)
		}
	}
}
