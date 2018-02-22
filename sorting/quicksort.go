package sorting

func Quicksort(arr []int) (res []int) {
	// check if empty (base case)
	if len(arr) <= 1 {
		return arr
	}
	// choose pivot value and create empty arrays
	pivotIndex := int((len(arr) -1)/2)
	left := []int{}
	right := []int{}
	for i, term := range arr {
		// exclude the pivot
		if i != pivotIndex {
			// move all less or equal into left
			if term <= arr[pivotIndex] {
				left = append(left, term)
			} else {
				// move all more into right
				right = append(right, term)
			}
		}
	}
	// recursively call append(quicksort(left), quicksort(right)...)
	return append(Quicksort(left), append([]int{arr[pivotIndex]}, Quicksort(right)...)...)
}
