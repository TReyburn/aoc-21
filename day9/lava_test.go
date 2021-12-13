package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLowPoints(t *testing.T) {
	wantLen := 4
	want := []int{1, 0, 5, 5}

	got := FindLowPoints(testData)

	assert.Equal(t, wantLen, len(got))
	assert.Equal(t, want, got)
}

func TestCalculateRisk(t *testing.T) {
	want := 15

	got := CalculateRisk(FindLowPoints(testData))

	assert.Equal(t, want, got)
}
