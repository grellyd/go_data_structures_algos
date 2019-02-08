package main

import "testing"

var tests = []struct{
	input string
	output bool
}{
	{"", true},
	{"a", true},
	{"aa", false},
	{"ab", true},
	{"abba", false},
	{"continue", false},
	{"failure", true},
}

func TestIsUniqueRange(t *testing.T) {
	for _, test := range tests {
		if res := IsUniqueRange(test.input); res != test.output {
			t.Errorf("IsUniqueRange(%q) = %v", test.input, res)
		}
	}
}

func TestIsUniqueDecoding(t *testing.T) {
	for _, test := range tests {
		if res := IsUniqueDecoding(test.input); res != test.output {
			t.Errorf("IsUniqueDecoding(%q) = %v", test.input, res)
		}
	}
}

func TestIsUniqueDecoded(t *testing.T) {
	for _, test := range tests {
		if res := IsUniqueDecoded(test.input); res != test.output {
			t.Errorf("IsUniqueDecoded(%q) = %v", test.input, res)
		}
	}
}

