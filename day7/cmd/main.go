package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day7"
)

func main() {
	pos := day7.AlignByMedian(day7.Day7Data)
	fuelCost := day7.CalculateFuelCost(day7.Day7Data, pos)

	fmt.Println("Day 7: Part 1: ", fuelCost)

	pos2 := day7.AlignByMean(day7.Day7Data)
	fuelCost2 := day7.CalculateAcceleratingFuelCosts(day7.Day7Data, pos2-1)

	fmt.Println("Day 7: Part 2: ", fuelCost2)
}
