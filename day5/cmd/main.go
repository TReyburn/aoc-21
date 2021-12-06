package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day5"
)

func main() {
	grid := day5.NewGrid()
	err := grid.BulkAddLines(day5.DayFiveData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 5: Part 1: ", grid.CountOverlaps())

	grid2 := day5.NewGrid()
	grid2.EnableDiagonals()
	err = grid2.BulkAddLines(day5.DayFiveData)
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 5: Part 2: ", grid2.CountOverlaps())
}
