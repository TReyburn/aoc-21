package main

import (
	"fmt"
	"github.com/TReyburn/aoc-21/day2"
)

func main() {
	fmt.Println("Day 2: Part 1: ", day2.DriveOutcome(day2.DriveSub(day2.DayTwoData)))
	fmt.Println("Day 2: Part 2: ", day2.DriveOutcome(day2.DriveSubWithAim(day2.DayTwoData)))
}
