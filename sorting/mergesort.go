package sorting

func Mergesort(data []int) (result [] int) {
	// base case: Single size list
	if len(data) <= 1 {
		return data
	}

	// recursive case: break into two equally sized lists
	left := data[:(len(data)/2)]
	right := data[(len(data)/2):]
	// merge the results of calling mergesort on each half
	return merge(Mergesort(left), Mergesort(right))
}

func merge(left []int, right []int) (combined_sorted []int) {
	combined_sorted = make([]int, len(left) + len(right))
	li := 0 
	ri := 0 
	for i := range combined_sorted {
		if li == len(left) {
			combined_sorted[i] = right[ri]
			ri += 1
		} else if ri == len(right) {
			combined_sorted[i] = left[li]
			li += 1
		} else if left[li] <= right[ri] {
			combined_sorted[i] = left[li]
			li += 1
		} else {
			combined_sorted[i] = right[ri]
			ri += 1
		}
	}
	return combined_sorted
}
