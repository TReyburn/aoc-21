package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day3"
)

func main() {
	fmt.Println("Day 3: Part 1: ", day3.CalculatePower(day3.RunPowerDiagnostic(day3.DayThreeData)))
	fmt.Println("Day 3: Part 2: ", day3.CalculatePower(day3.FindLifeSupportRatings(day3.DayThreeData)))
}
