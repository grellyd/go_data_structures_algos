package fib

var FibBase = []int{0, 1} // 1

func ComputeNextValue(values []int) int { 
	return values[len(values)-1] + values[len(values)-2] // 1 * O(len([]))
}

func AddNextValue(values []int) []int {
	return append(values, ComputeNextValue(values)) // 1 * O(append)
}

func GetFibs(num int) []int {
	if num <= 2 {   // 1
		return []int{0, 1} // 1
	}
	values := FibBase // 1
	for n := 0; n < num-2; n++ {   //  n-2 --> n
		values = AddNextValue(values) // 1 * O(add...)
	}
	return values // 1
}

func IsFib(value int) bool {
	if value < 0 {
		return false
	}
	if value == 0 || value == 1 {
		return true
	}
	values := FibBase
	for { //  O(value)  --> O(n)
		newestValue := ComputeNextValue(values) // 1 * O(compute..)
		switch {
		case newestValue == value:  // 1
			return true // 1
		case newestValue > value: // 1
			return false // 1
		default:
			values = append(values, newestValue) // 1 * O(append)
		}
	}
}

func IsFibs(values []int) bool {
	areFibs := true
	for _, value := range values { // O(value) --> O(n)
		for n:=0; n< len(values); n++ {
			// do something in constant time
		}
		areFibs = areFibs && IsFib(value)  // 1
	}
	return areFibs
}

func NthFib(n int) int {
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return nthFibHelper(FibBase, n-2)
}

func nthFibHelper(values []int, n int) int {
	if n == 0 {
		return values[len(values)-1]
	}
	return nthFibHelper(AddNextValue(values), n-1)
}
