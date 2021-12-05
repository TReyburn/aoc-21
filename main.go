package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day1"
	"github.com/TReyburn/aoc-21/day2"
)

var (
	d1Start, _ = day1.DoublyLinkedList(day1.PartOne)
)

func main() {
	fmt.Println("Day 1: Part 1: ", day1.ScanIncreases(d1Start))
	fmt.Println("Day 1: Part 2: ", day1.ScanExtendedIncreases(d1Start))

	fmt.Println("Day 2: Part 1: ", day2.DriveOutcome(day2.DriveSub(day2.DayTwoData)))
	fmt.Println("Day 2: Part 2: ", day2.DriveOutcome(day2.DriveSubWithAim(day2.DayTwoData)))
}
