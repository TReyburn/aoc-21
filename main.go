package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day1"
	"github.com/TReyburn/aoc-21/day2"
	"github.com/TReyburn/aoc-21/day3"
	"github.com/TReyburn/aoc-21/day4"
)

var (
	dayOne   = false
	dayTwo   = false
	dayThree = false
	dayFour  = true
)

func main() {
	if dayOne {
		d1Start, _ := day1.DoublyLinkedList(day1.PartOne)
		fmt.Println("Day 1: Part 1: ", day1.ScanIncreases(d1Start))
		fmt.Println("Day 1: Part 2: ", day1.ScanExtendedIncreases(d1Start))
	}

	if dayTwo {
		fmt.Println("Day 2: Part 1: ", day2.DriveOutcome(day2.DriveSub(day2.DayTwoData)))
		fmt.Println("Day 2: Part 2: ", day2.DriveOutcome(day2.DriveSubWithAim(day2.DayTwoData)))
	}

	if dayThree {
		fmt.Println("Day 3: Part 1: ", day3.CalculatePower(day3.RunPowerDiagnostic(day3.DayThreeData)))
		fmt.Println("Day 3: Part 2: ", day3.CalculatePower(day3.FindLifeSupportRatings(day3.DayThreeData)))
	}

	if dayFour {
		bingo := day4.LoadBingo(day4.DayFourData)
		fmt.Println("Day 4: Part 1: ", day4.CalculateScore(bingo.RunGame()))
		newBingo := day4.LoadBingo(day4.DayFourData)
		fmt.Println("Day 4: Part 2: ", day4.CalculateScore(newBingo.RunGameUntilLastBoardWins()))
	}
}
