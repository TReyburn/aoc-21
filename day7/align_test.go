package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlignByMedian(t *testing.T) {
	want := 2

	got := AlignByMedian(testData)
	assert.Equal(t, want, got)
}

func TestCalculateFuelCost(t *testing.T) {
	var costTest = []struct {
		want int
		pos  int
	}{{37, 2}, {41, 1}, {39, 3}, {71, 10}}

	for _, test := range costTest {
		fc := test
		got := CalculateFuelCost(testData, fc.pos)
		assert.Equal(t, fc.want, got)
	}
}

func TestAlignByMean(t *testing.T) {
	want := 5

	got := AlignByMean(testData)
	assert.Equal(t, want, got)
}

func TestCalculateAcceleratingFuelCosts(t *testing.T) {
	var costTest = []struct {
		want int
		pos  int
	}{{206, 2}, {168, 5}}

	for _, tc := range costTest {
		got := CalculateAcceleratingFuelCosts(testData, tc.pos)
		assert.Equal(t, tc.want, got)
	}
}

func TestCalculateAcceleratingFuelCosts2(t *testing.T) {
	var costTest = []struct {
		want int
		pos  int
		data []int
	}{
		{66, 5, []int{16}},
		{10, 5, []int{1}},
		{6, 5, []int{2}},
		{15, 5, []int{0}},
		{1, 5, []int{4}},
		{3, 5, []int{7}},
		{45, 5, []int{14}},
	}

	for _, tc := range costTest {
		got := CalculateAcceleratingFuelCosts(tc.data, tc.pos)
		assert.Equal(t, tc.want, got)
	}
}
