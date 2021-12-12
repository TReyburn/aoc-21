package day7

import "github.com/TReyburn/aoc-21/algorithms"

func AlignByMedian(d []int) int {
	sorted := algorithms.QuickSort(d)
	return algorithms.Median(sorted)
}

func CalculateFuelCost(input []int, pos int) int {
	var total int
	for _, i := range input {
		total += algorithms.Max(i, pos) - algorithms.Min(i, pos)
	}
	return total
}

func AlignByMean(d []int) int {
	return algorithms.Mean(d)
}

func CalculateAcceleratingFuelCosts(input []int, pos int) int {
	var total int
	for _, i := range input {
		steps := algorithms.Max(i, pos) - algorithms.Min(i, pos)
		for x := 1; x <= steps; x++ {
			total += x
		}
	}
	return total
}
