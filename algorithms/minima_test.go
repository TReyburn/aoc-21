package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	var minTests = []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1},
		{1, 1, 1},
		{0, -1, -1},
		{2, 1, 1},
	}

	for _, tc := range minTests {
		got := Min(tc.a, tc.b)
		assert.Equal(t, tc.want, got)
	}
}

func TestMax(t *testing.T) {
	var maxTests = []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 2},
		{1, 1, 1},
		{0, -1, 0},
		{2, 1, 2},
	}
	
	for _, tc := range maxTests {
		got := Max(tc.a, tc.b)
		assert.Equal(t, tc.want, got)
	}
}