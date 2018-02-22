package sorting

import "fmt"

func Bubblesort(lst []int) []int {
	fmt.Printf("Bubblesort: %v --> ", lst) 
	// for num
	for j := 0; j < len(lst); j++ {
		curr := 0
		for i := 1; i < len(lst); i++ {
			item := lst[curr]
			if item > lst[i] {
				lst[i-1] = lst[i]
				lst[i] = item
			}
			curr = i
		}
	
	}
	fmt.Printf("%v\n", lst) 
	return lst
}
