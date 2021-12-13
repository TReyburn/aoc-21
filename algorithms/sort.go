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

func QuickSort32(arr []int32) []int32 {
	if len(arr) < 2 {
		return arr
	}
	less := make([]int32, 0)
	greater := make([]int32, 0)
	result := make([]int32, 0)
	pivot := arr[0]

	for _, i := range arr[1:] {
		if i <= pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}

	result = append(result, QuickSort32(less)...)
	result = append(result, pivot)
	result = append(result, QuickSort32(greater)...)
	return result
}

func SortString(s string) string {
	runes := []rune(s)
	res := QuickSort32(runes)
	return string(res)
}
