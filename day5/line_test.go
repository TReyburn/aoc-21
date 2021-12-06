package day5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPoint(t *testing.T) {
	wantX := 2
	wantY := 4
	wantStr := fmt.Sprintf("%v,%v", wantX, wantY)

	p, err := NewPoint(wantStr)
	assert.NoError(t, err)

	assert.Equal(t, wantX, p.x)
	assert.Equal(t, wantY, p.y)
	assert.Equal(t, wantStr, p.str)
}

func TestGrid_CreatePoint(t *testing.T) {
	wantX := 2
	wantY := 4
	wantPointsNum := 1
	wantPointCount := 1

	grid := NewGrid()
	p, err := NewPoint(fmt.Sprintf("%v,%v", wantX, wantY))
	assert.NoError(t, err)

	grid.AddPoint(p)

	assert.NoError(t, err)
	assert.Equal(t, wantPointsNum, len(grid.points))
	assert.Equal(t, wantPointCount, grid.counts[p])
}

func TestGrid_AddLine_VerticalPos(t *testing.T) {
	line := `1,1 -> 1,3`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_HorizontalNeg(t *testing.T) {
	line := `9,7 -> 7,7`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_Diagonal_NonEnabled(t *testing.T) {
	// This test is to ensure that diagonals don't get added by default
	line := `9,7 -> 7,9`
	wantPointsNum := 0
	wantPointCount := 0

	grid := NewGrid()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_Diagonal_Enabled(t *testing.T) {
	line := `1,1 -> 3,3`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()
	grid.EnableDiagonals()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_Diagonal_Enabled2(t *testing.T) {
	line := `9,7 -> 7,9`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()
	grid.EnableDiagonals()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_Diagonal_Enabled3(t *testing.T) {
	line := `9,7 -> 7,5`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()
	grid.EnableDiagonals()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_AddLine_Diagonal_Enabled4(t *testing.T) {
	line := `7,9 -> 9,7`
	wantPointsNum := 3
	wantPointCount := 1

	grid := NewGrid()
	grid.EnableDiagonals()

	err := grid.AddLine(line)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_BulkAddLines(t *testing.T) {
	lines := `0,0 -> 0,8
2,2 -> 2,1`

	wantPointsNum := 11
	wantPointCount := 1

	grid := NewGrid()

	err := grid.BulkAddLines(lines)
	assert.NoError(t, err)

	assert.Equal(t, wantPointsNum, len(grid.points))
	for _, v := range grid.counts {
		assert.Equal(t, wantPointCount, v)
	}
}

func TestGrid_CountOverlaps(t *testing.T) {
	wantOverlapCount := 5

	grid := NewGrid()

	err := grid.BulkAddLines(testData)
	assert.NoError(t, err)

	gotOverlapCount := grid.CountOverlaps()

	assert.Equal(t, wantOverlapCount, gotOverlapCount)
}

func TestGrid_CountOverlaps_EdgeCase(t *testing.T) {
	wantOverlapCount := 51
	lines := `10,100 -> 120,100
70,100 -> 500,100`

	grid := NewGrid()

	err := grid.BulkAddLines(lines)

	assert.NoError(t, err)

	gotOverlapCount := grid.CountOverlaps()
	assert.Equal(t, wantOverlapCount, gotOverlapCount)
}

func TestGrid_CountOverlaps_EdgeCase2(t *testing.T) {
	wantOverlapCount := 1
	lines := `579,680 -> 581,680
580,679 -> 580,681`

	grid := NewGrid()

	err := grid.BulkAddLines(lines)

	assert.NoError(t, err)

	gotOverlapCount := grid.CountOverlaps()
	assert.Equal(t, wantOverlapCount, gotOverlapCount)
}

func TestGrid_CountOverlaps_Diagonal(t *testing.T) {
	wantOverlapCount := 12

	grid := NewGrid()
	grid.EnableDiagonals()
	err := grid.BulkAddLines(testData)
	assert.NoError(t, err)

	gotOverlapCount := grid.CountOverlaps()
	assert.Equal(t, wantOverlapCount, gotOverlapCount)
}

func TestGrid_CountOverlaps_Diagonal2(t *testing.T) {
	wantOverlapCount := 1
	lines := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4`

	grid := NewGrid()
	grid.EnableDiagonals()
	err := grid.BulkAddLines(lines)
	assert.NoError(t, err)

	gotOverlapCount := grid.CountOverlaps()
	assert.Equal(t, wantOverlapCount, gotOverlapCount)
}
