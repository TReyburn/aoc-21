package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	start, end := DoublyLinkedList(testDataPartOne)

	assert.Equal(t, 199, start.Value())
	assert.Nil(t, start.Prev())
	assert.Equal(t, 200, start.Next().Value())

	assert.Equal(t, 263, end.Value())
	assert.Nil(t, end.Next())
	assert.Equal(t, 260, end.Prev().Value())
}
