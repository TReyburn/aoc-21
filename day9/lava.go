package day9

func FindLowPoints(input [][]int) []int {
	lowPoints := make([]int, 0)

	// nav points
	up := -1
	down := 1
	left := -1
	right := 1

	// bounds
	maxDown := len(input)
	maxRight := len(input[0])

	for rowIdx, row := range input {
		for colIdx, x := range row {
			lower := 0
			if up >= 0 {
				if x < input[up][colIdx] {
					lower++
				}
			} else {
				lower++
			}
			if down < maxDown && down > rowIdx {
				if x < input[down][colIdx] {
					lower++
				}
			} else {
				lower++
			}
			if left >= 0 {
				if x < input[rowIdx][left] {
					lower++
				}
			} else {
				lower++
			}
			if right < maxRight && right > colIdx {
				if x < input[rowIdx][right] {
					lower++
				}
			} else {
				lower++
			}
			if lower == 4 {
				lowPoints = append(lowPoints, x)
			}
			right++
			left++
		}
		right = 1
		left = -1
		up++
		down++
	}

	return lowPoints
}

func CalculateRisk(points []int) int {
	var riskTotal int

	for _, num := range points {
		riskTotal += 1 + num
	}

	return riskTotal
}
