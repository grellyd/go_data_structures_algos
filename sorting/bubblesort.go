package sorting

func Bubblesort(lst []int) []int {
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
	return lst
}
