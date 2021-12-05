package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadBoard(t *testing.T) {
	wantLength := 5
	wantRowLength := 5
	wantFirstRow := []int{22, 13, 17, 11, 0}
	wantLastRow := []int{1, 12, 20, 15, 19}
	input := `22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`

	output := LoadBoard(input)

	assert.Equal(t, wantLength, len(output.nums))
	for _, row := range output.nums {
		assert.Equal(t, wantRowLength, len(row))
	}
	assert.Equal(t, wantFirstRow, output.nums[0])
	assert.Equal(t, wantLastRow, output.nums[4])
}

func TestLoadBingo(t *testing.T) {
	wantDraw := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	wantBoardsLen := 3
	wantFirstBoard := LoadBoard(`22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`)
	wantLastBoard := LoadBoard(`14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`)

	bingo := LoadBingo(testData)

	assert.Equal(t, wantDraw, bingo.draw)
	assert.Equal(t, wantBoardsLen, len(bingo.boards))
	assert.Equal(t, wantFirstBoard, bingo.boards[0])
	assert.Equal(t, wantLastBoard, bingo.boards[2])
}

func TestBoard_Check(t *testing.T) {
	input := `22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`

	board := LoadBoard(input)

	assert.True(t, board.Check(22))
	assert.False(t, board.Check(32))
}

func TestBoard_Update(t *testing.T) {
	wantMarkedCount := 1
	input := `22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`

	board := LoadBoard(input)

	board.Update(9)

	assert.Equal(t, marked, board.nums[2][1])

	var gotCountMarked int
	for _, row := range board.nums {
		for _, num := range row {
			if num == marked {
				gotCountMarked++
			}
		}
	}

	assert.Equal(t, wantMarkedCount, gotCountMarked)
}

func TestBoard_Update_NoUpdate(t *testing.T) {
	wantMarkedCount := 0
	input := `22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`

	board := LoadBoard(input)

	board.Update(69)

	var gotCountMarked int
	for _, row := range board.nums {
		for _, num := range row {
			if num == marked {
				gotCountMarked++
			}
		}
	}

	assert.Equal(t, wantMarkedCount, gotCountMarked)
}

func TestBoard_IsBingo_NoBingo(t *testing.T) {
	input := `22 13 -1 11  0
8  2 23  4 24
21  -1 -1 16  7
6 10  3 18  5
1 12 20 -1 19`

	board := LoadBoard(input)

	assert.False(t, board.IsBingo())
}

func TestBoard_IsBingo_VerticalBingo(t *testing.T) {
	input := `22 -1 17 11  0
8  -1 23  4 24
21  -1 14 16  7
6 -1  3 18  5
1 -1 20 15 19`

	board := LoadBoard(input)

	assert.True(t, board.IsBingo())
}

func TestBoard_IsBingo_DownwardDiagonalBingo(t *testing.T) {
	input := `-1 13 17 11  0
8  -1 23  4 24
21  9 -1 16  7
6 10  3 -1  5
1 12 20 15 -1`

	board := LoadBoard(input)

	assert.False(t, board.IsBingo())
}

func TestBoard_IsBingo_UpwardDiagonalBingo(t *testing.T) {
	input := `22 13 17 11  -1
8  2 23  -1 24
21  9 -1 16  7
6 -1  3 18  5
-1 12 20 15 19`

	board := LoadBoard(input)

	assert.False(t, board.IsBingo())
}

func TestBoard_IsDiagonalBingo_DownwardBingo(t *testing.T) {
	input := `-1 13 17 11  0
8  -1 23  4 24
21  9 -1 16  7
6 10  3 -1  5
1 12 20 15 -1`

	board := LoadBoard(input)

	assert.True(t, board.IsDiagonalBingo())
}

func TestBoard_IsDiagonalBingo_UpwardBingo(t *testing.T) {
	input := `22 13 17 11  -1
8  2 23  -1 24
21  9 -1 16  7
6 -1  3 18  5
-1 12 20 15 19`

	board := LoadBoard(input)

	assert.True(t, board.IsDiagonalBingo())
}

func TestBoard_SumUnmarked(t *testing.T) {
	wantUnmarkedSum := 188
	input := `-1 -1 -1 -1  -1
10 16 15  -1 19
18  8 -1 26 20
22 -1 13  6  -1
-1  -1 12  3  -1`

	board := LoadBoard(input)

	assert.Equal(t, wantUnmarkedSum, board.SumUnmarked())
}

func TestBingo_RunGame(t *testing.T) {
	wantWinningNum := 24
	wantSumUnmarked := 188

	bingo := LoadBingo(testData)

	res := bingo.RunGame()

	assert.Equal(t, wantWinningNum, res.WinningNum)
	assert.Equal(t, wantSumUnmarked, res.SumUnmarked)
}

func TestCalculateScore(t *testing.T) {
	wantScore := 4512

	bingo := LoadBingo(testData)

	res := bingo.RunGame()

	gotScore := CalculateScore(res)

	assert.Equal(t, wantScore, gotScore)
}

func TestBingo_RunGameUntilLastBoardWins(t *testing.T) {
	wantSumUnmarked := 148
	wantWinningNum := 13
	wantScore := 1924

	bingo := LoadBingo(testData)

	res := bingo.RunGameUntilLastBoardWins()

	assert.Equal(t, wantSumUnmarked, res.SumUnmarked)
	assert.Equal(t, wantWinningNum, res.WinningNum)

	gotScore := CalculateScore(res)

	assert.Equal(t, wantScore, gotScore)
}
