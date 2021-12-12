package algorithms

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	less := make([]int, 0)
	greater := make([]int, 0)
	result := make([]int, 0)
	pivot := arr[0]

	for _, i := range arr[1:] {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}

	result = append(result, QuickSort(less)...)
	result = append(result, pivot)
	result = append(result, QuickSort(greater)...)
	return result
}
