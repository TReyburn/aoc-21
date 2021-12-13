package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day8"
)

func main() {
	s1 := day8.NewScreen()
	search := []int{1, 4, 7, 8}

	inputs, outputs := day8.ParseData(day8.Day8Data)
	for _, out := range outputs {
		s1.Process(out)
	}

	count := s1.CountNums(search)

	fmt.Println("Day 8: Part 1: ", count)

	var totalCount int

	for idx, out := range outputs {
		s2 := day8.NewScreen()
		s2.LoadInstructions(inputs[idx])
		totalCount += s2.ProcessFromInstructions(out)
	}

	fmt.Println("Day 8: Part 2: ", totalCount)
}
