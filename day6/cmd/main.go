package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day6"
)

func main() {
	s := day6.NewSchool(day6.DaySixData)
	s.ProcessDays(80)
	fmt.Println("Day 6: Part 1: ", s.CountFish())

	s2 := day6.NewSchool(day6.DaySixData)
	s2.ProcessDays(256)
	fmt.Println("Day 6: Part 2: ", s2.CountFish())
}
