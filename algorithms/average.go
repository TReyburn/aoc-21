package algorithms

import "math"

// Median expects a sorted list. For a list of even length, Median will round the position
func Median(d []int) int {
	l := len(d)
	return d[l/2]
}

func Mean(d []int) int {
	var sum int
	for _, i := range d {
		sum += i
	}
	return int(math.Round(float64(sum) / float64(len(d))))
}
