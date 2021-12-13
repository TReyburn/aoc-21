package day8

import (
	"fmt"
	"github.com/TReyburn/aoc-21/algorithms"
)

type Number struct {
	count int
}

type Screen struct {
	nums          map[int]*Number
	lookup        map[string]int
	reverseLookup map[int]string
}

func (s *Screen) Process(output []string) {
	for _, n := range output {
		i := len(n)
		switch i {
		case 2: // case for One
			one, ok := s.nums[1]
			if !ok {
				one = &Number{}
				s.nums[1] = one
			}
			one.count++

		case 3: // case for Seven
			seven, ok := s.nums[7]
			if !ok {
				seven = &Number{}
				s.nums[7] = seven
			}
			seven.count++

		case 4: // case for Four
			four, ok := s.nums[4]
			if !ok {
				four = &Number{}
				s.nums[4] = four
			}
			four.count++

		case 5: // case for Two, Three, Five,
		case 6: // case for Zero, Six, Nine
		case 7: // case for Eight
			eight, ok := s.nums[8]
			if !ok {
				eight = &Number{}
				s.nums[8] = eight
			}
			eight.count++
		}
	}
}

func (s *Screen) LoadInstructions(input []string) {
	fives := make([]string, 0)
	sixes := make([]string, 0)

	// lazy load known numbers
	for _, n := range input {
		i := len(n)
		switch i {
		case 2:
			s.lookup[algorithms.SortString(n)] = 1
			s.reverseLookup[1] = algorithms.SortString(n)
		case 3:
			s.lookup[algorithms.SortString(n)] = 7
			s.reverseLookup[7] = algorithms.SortString(n)
		case 4:
			s.lookup[algorithms.SortString(n)] = 4
			s.reverseLookup[4] = algorithms.SortString(n)
		case 5:
			fives = append(fives, n)
		case 6:
			sixes = append(sixes, n)
		case 7:
			s.lookup[algorithms.SortString(n)] = 8
			s.reverseLookup[8] = algorithms.SortString(n)
		}
	}

	for _, f := range fives {
		x := algorithms.StringOverlap(f, s.reverseLookup[4])
		if x == 2 {
			s.lookup[algorithms.SortString(f)] = 2
			s.reverseLookup[2] = algorithms.SortString(f)
		} else {
			z := algorithms.StringOverlap(f, s.reverseLookup[1])
			if z == 1 {
				s.lookup[algorithms.SortString(f)] = 5
				s.reverseLookup[5] = algorithms.SortString(f)
			}
		}
		y := algorithms.StringOverlap(f, s.reverseLookup[1])
		if y == 2 {
			s.lookup[algorithms.SortString(f)] = 3
			s.reverseLookup[3] = algorithms.SortString(f)
		}
	}

	for _, six := range sixes {
		x := algorithms.StringOverlap(six, s.reverseLookup[7])
		if x == 2 {
			s.lookup[algorithms.SortString(six)] = 6
			s.reverseLookup[6] = algorithms.SortString(six)
		} else {
			z := algorithms.StringOverlap(six, s.reverseLookup[1])
			if z == 2 {
				s.lookup[algorithms.SortString(six)] = 0
				s.reverseLookup[0] = algorithms.SortString(six)
			}
		}
		y := algorithms.StringOverlap(six, s.reverseLookup[4])
		if y == 4 {
			s.lookup[algorithms.SortString(six)] = 9
			s.reverseLookup[9] = algorithms.SortString(six)
		}
	}
}

func (s *Screen) ProcessFromInstructions(output []string) int {
	var processedNum int
	pos := 1000

	for _, out := range output {
		num, ok := s.lookup[algorithms.SortString(out)]
		if !ok {
			fmt.Println("Unrecognized string: ", out)
		} else {
			processedNum += num * pos
		}
		pos /= 10
	}
	return processedNum
}

func (s *Screen) CountNums(ns []int) int {
	var total int
	for _, n := range ns {
		num, ok := s.nums[n]
		if ok {
			total += num.count
		}
	}
	return total
}

func NewScreen() *Screen {
	m := make(map[int]*Number, 0)
	l := make(map[string]int, 0)
	r := make(map[int]string, 0)
	return &Screen{nums: m, lookup: l, reverseLookup: r}
}

func Overlap(a string, b string) int {
	return 0
}
