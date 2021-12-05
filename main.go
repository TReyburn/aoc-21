package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day1"
	"github.com/TReyburn/aoc-21/day2"
	"github.com/TReyburn/aoc-21/day3"
)

var (
	dayOne   = false
	dayTwo   = false
	dayThree = true
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
		fmt.Println("day 3: Part 1: ", day3.CalculatePower(day3.RunPowerDiagnostic(day3.DayThreeData)))
		fmt.Println("day 3: Part 2: ", day3.CalculatePower(day3.FindLifeSupportRatings(day3.DayThreeData)))
	}
}
