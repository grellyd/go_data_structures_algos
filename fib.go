package fib

var FibBase = []int{0, 1}

func ComputeNextValue(values []int) int {
	return values[len(values) - 1] + values[len(values) - 2]
}

func AddNextValue(values []int) []int {
	return append(values, ComputeNextValue(values))
}

func GetFibs(num int) []int {
	if num <= 2 { return []int{0, 1} }
	values := FibBase
	for n := 0; n < num - 2; n++ {
		values = AddNextValue(values)
	}
	return values
}

func IsFib(value int) bool {
	if value < 0 { return false }
	if value == 0 || value == 1 { return true }
	values := FibBase
	for {
		newestValue := ComputeNextValue(values)
		switch {
		case newestValue == value:
			return true
		case newestValue > value:
			return false
		default:	
			values = append(values, newestValue)
		}
	}
}

func IsFibs(values []int) bool {
	areFibs := true
	for _, value := range values {
		areFibs = areFibs && IsFib(value)
	}
	return areFibs
}
		
func NthFib(n int) int {
	if n <= 1 { return 0 }
	if n == 2 { return 1 }
	return nthFibHelper(FibBase, n - 2)
}

func nthFibHelper(values []int, n int) int {
	if n == 0 {
		return values[len(values) - 1]
	}
	return nthFibHelper(AddNextValue(values), n - 1)
}
