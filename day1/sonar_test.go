package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanIncreases(t *testing.T) {
	want := 7
	start, _ := DoublyLinkedList(testDataPartOne)

	got := ScanIncreases(start)

	assert.Equal(t, want, got)
}

func TestScanExtendedIncreases(t *testing.T) {
	want := 5
	start, _ := DoublyLinkedList(testDataPartOne)

	got := ScanExtendedIncreases(start)

	assert.Equal(t, want, got)
}
