package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day9"
)

func main() {
	lowPoints := day9.FindLowPoints(day9.Day9Data)
	risk := day9.CalculateRisk(lowPoints)
	fmt.Println("Day 1: Part 1: ", risk)
}
