package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var sortTest = []struct {
		data []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 4, 3}, []int{1, 2, 3, 4}},
		{[]int{2, 2, 4, 3, 1}, []int{1, 2, 2, 3, 4}},
	}

	for _, test := range sortTest {
		got := QuickSort(test.data)
		assert.Equal(t, test.want, got)
	}

}
