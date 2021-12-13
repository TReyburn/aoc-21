package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedian(t *testing.T) {
	var medianTests = []struct {
		data []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{0, 1, 1, 2, 2, 2, 2, 4, 7, 14, 16}, 2},
	}

	for _, tc := range medianTests {
		got := Median(tc.data)
		assert.Equal(t, tc.want, got)
	}
}

func TestMean(t *testing.T) {
	var meanTests = []struct {
		data []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{0, 1, 1, 2, 2, 2, 2, 4, 7, 14, 16}, 5},
	}

	for _, tc := range meanTests {
		got := Mean(tc.data)
		assert.Equal(t, tc.want, got)
	}
}
