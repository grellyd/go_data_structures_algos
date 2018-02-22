package main

import (
	"fmt"
	"strconv"
)

// Given:

func main() {
	fmt.Println(bar(
		[][]string{
			{"false", "C"},
			{"false", "S"},
			{"true", "K"},
			{"true", "K"},
			{"false", "F"},
			{"false", "B"},
			{"true", "K"},
			{"false", "Z"},
			{"false", "T"},
			{"true", "E"},
			{"false", "S"},
			{"false", "U"},
			{"false", "R"},
			{"true", "N"},
			{"false", "Z"},
			{"false", "U"},
			{"true", "N"},
			{"false", "S"},
			{"false", "D"},
			{"true", "Y"},
			{"false", "J"},
			{"false", "K"},
			{"false", "T"},
		},
	))
}

func bar(z [][]string) string {
	output := ""
	for i, row := range z {
		parsedBool, _ := strconv.ParseBool(row[0])
		if foo(i, parsedBool) {
			output += row[1]
		}
	}
	return output
}

func foo(a int, b bool) bool {
	res := false
	if a > 5 {
		res = (b || res) && !(b && res)
	}
	return res
}

// Produce:

// Function main
// Entry point for the go program
// Arguments: None
// Returns: Nothing
func main_annoted() {
	fmt.Println(bar( // call to bar wrapped with a call to Println from the fmt package
		[][]string{ // annoymous instantiation of a multidimensional string array
			{"false", "C"},
			{"false", "S"},
			{"true", "K"},
			{"true", "K"},
			{"false", "F"},
			{"false", "B"},
			{"true", "K"},
			{"false", "Z"},
			{"false", "T"},
			{"true", "E"},
			{"false", "S"},
			{"false", "U"},
			{"false", "R"},
			{"true", "N"},
			{"false", "Z"},
			{"false", "U"},
			{"true", "N"},
			{"false", "S"},
			{"false", "D"},
			{"true", "Y"},
			{"false", "J"},
			{"false", "K"},
			{"false", "T"},
		},
	))
}

// Function bar
// Combines any strings which statisfy foo
// Arguments:
//  - Multidimensional String Array z
// Returns:
//  - String
func bar_annotated(z [][]string) string {
	output := ""            // declaration of output variable
	for i, row := range z { // Control Flow: for loop with indicies for each value in z
		parsedBool, _ := strconv.ParseBool(row[0]) // parsing the first item in the row into a boolean from a string using the strconv package
		if foo(i, parsedBool) {                    // Control Flow: if where the conditional is the call to foo
			output += row[1] // string concatination
		}
	}
	return output
}

// Function foo
// Returns true if a is greater than 5 and b is false.
// Arguments:
// 	- Integer a
//  - Boolean b
// Returns:
//  - Boolean
func foo_annotated(a int, b bool) bool {
	res := false // intermediate variable assignment of res to false
	if a > 5 {   // Control Flow: If statement where the conditional is (a > 5)
		res = (b || res) && !(b && res) // assignment to res of the XOR of b and false. Will only be true when b is true, as res is always false
	}
	return res
}

// Be sure to identifiy what you can in terms of:
//  - control flow structures
//  - data structures
//  - function contracts (names, arguments, returns, purpose)
//  - types
//  - iteration and recursion
//  - function calls
