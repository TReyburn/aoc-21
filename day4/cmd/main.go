package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day4"
)

func main() {
	bingo := day4.LoadBingo(day4.DayFourData)
	fmt.Println("Day 4: Part 1: ", day4.CalculateScore(bingo.RunGame()))
	newBingo := day4.LoadBingo(day4.DayFourData)
	fmt.Println("Day 4: Part 2: ", day4.CalculateScore(newBingo.RunGameUntilLastBoardWins()))
}
