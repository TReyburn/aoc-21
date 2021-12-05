package day4

import (
	"fmt"
	"strconv"
	"strings"
)

const marked = -1

type GameResult struct {
	WinningNum   int
	SumUnmarked  int
	winningBoard *Board
}

type Bingo struct {
	draw   []int
	boards []*Board
}

func (b *Bingo) RunGame() GameResult {
	var res GameResult

	for _, num := range b.draw {
		for _, board := range b.boards {
			board.Update(num)
			if board.IsBingo() {
				res.WinningNum = num
				res.SumUnmarked = board.SumUnmarked()
				res.winningBoard = board
				return res
			}
		}
	}

	return res
}

func (b *Bingo) RunGameUntilLastBoardWins() GameResult {
	var res GameResult
	var s struct{}
	var countWinningBoards int
	winningBoards := make(map[*Board]struct{})

	totalBoards := len(b.boards)

	for _, num := range b.draw {
		for _, board := range b.boards {
			if _, ok := winningBoards[board]; !ok {
				board.Update(num)
				if board.IsBingo() {
					winningBoards[board] = s
					countWinningBoards++
					if countWinningBoards == totalBoards {
						res.WinningNum = num
						res.SumUnmarked = board.SumUnmarked()
						res.winningBoard = board
						return res
					}
				}
			}
		}
	}
	return res
}
func CalculateScore(res GameResult) int {
	return res.WinningNum * res.SumUnmarked
}

type Board struct {
	nums  [][]int
	check map[int]struct{}
}

func (b *Board) Check(x int) bool {
	_, ok := b.check[x]
	return ok
}

func (b *Board) Update(x int) {
	if b.Check(x) {
	rows:
		for rowIdx, row := range b.nums {
			for numIdx, num := range row {
				if num == x {
					b.nums[rowIdx][numIdx] = marked
					break rows
				}
			}
		}
	}
}

func (b *Board) IsBingo() bool {
	// O(2n^2)
	var bingo bool

	// check horizontal - O(n^2)
	for _, row := range b.nums {
		markedCount := 0
		for _, num := range row {
			if num != marked {
				break
			}
			markedCount++
		}
		if markedCount == len(b.nums[0]) {
			bingo = true
			return bingo
		}
	}

	// check vertical - O(n^2)
	for i := 0; i < len(b.nums[0]); i++ {
		markedCount := 0
		for _, row := range b.nums {
			if row[i] != marked {
				break
			}
			markedCount++
		}
		if markedCount == len(b.nums[0]) {
			bingo = true
			return bingo
		}
	}

	return bingo
}

func (b *Board) IsDiagonalBingo() bool {
	// O(n)
	var bingo bool

	// check diagonal - O(n)
	var upMarked, downMarked int
	for i := 0; i < len(b.nums[0]); i++ {
		if downMarked != i && upMarked != i {
			break
		}
		if downMarked == i {
			if b.nums[i][i] == marked {
				downMarked++
			}
		}
		if upMarked == i {
			if b.nums[len(b.nums[0])-i-1][i] == marked {
				upMarked++
			}
		}
	}
	if upMarked == len(b.nums[0]) || downMarked == len(b.nums[0]) {
		bingo = true
		return bingo
	}

	return bingo
}

func (b *Board) SumUnmarked() int {
	var unmarkedSum int

	for _, row := range b.nums {
		for _, num := range row {
			if num != marked {
				unmarkedSum += num
			}
		}
	}

	return unmarkedSum
}

func LoadBingo(input string) Bingo {
	var bingo Bingo

	bingoSplit := strings.SplitN(input, "\n", 3)

	drawSplit := strings.Split(bingoSplit[0], ",")
	bingo.draw = convertStringsToInts(drawSplit)

	boardSplit := strings.Split(bingoSplit[2], "\n")
	var board string
	for _, row := range boardSplit {
		if row == "" {
			bingo.boards = append(bingo.boards, LoadBoard(board))
			board = ""
		} else {
			board += fmt.Sprintf("%s\n", row)
		}
	}
	if board != "" {
		bingo.boards = append(bingo.boards, LoadBoard(board))
	}

	return bingo
}

func LoadBoard(input string) *Board {
	var board Board
	var s struct{}
	set := make(map[int]struct{})
	board.check = set
	boardSplit := strings.Split(input, "\n")

	for _, row := range boardSplit {
		if row != "" {
			rowSplit := strings.Split(row, " ")
			rowNums := convertStringsToInts(rowSplit)
			board.nums = append(board.nums, rowNums)
			for _, num := range rowNums {
				board.check[num] = s
			}
		}
	}
	return &board
}

func convertStringsToInts(s []string) []int {
	i := make([]int, 0)
	for _, sNum := range s {
		if sNum != "" {
			iNum, err := strconv.Atoi(sNum)
			if err != nil {
				fmt.Printf("Unexpected err: %v\n", err)
			}
			i = append(i, iNum)
		}
	}
	return i
}
