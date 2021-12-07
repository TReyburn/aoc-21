package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSchool(t *testing.T) {
	wantInitTimer := 8
	wantResetTimer := 6
	wantCount := 0

	s := NewSchool([]int{})

	assert.Equal(t, wantInitTimer, s.initTimer)
	assert.Equal(t, wantResetTimer, s.resetTimer)
	for _, v := range s.school {
		assert.Equal(t, wantCount, v)
	}
}

func TestSchool_Seed(t *testing.T) {
	input := []int{1, 2, 3}
	wantInitCounts := map[int]int{1: 1, 2: 1, 3: 1}
	seed1 := []int{1}
	wantSeed1Counts := map[int]int{1: 2, 2: 1, 3: 1}
	seed2 := []int{1, 2}
	wantSeed2Counts := map[int]int{1: 3, 2: 2, 3: 1}
	var seed3 []int

	s := NewSchool(input)
	for k, v := range wantInitCounts {
		assert.Equal(t, v, s.school[k])
	}

	s.Seed(seed1...)
	for k, v := range wantSeed1Counts {
		assert.Equal(t, v, s.school[k])
	}

	s.Seed(seed2...)
	for k, v := range wantSeed2Counts {
		assert.Equal(t, v, s.school[k])
	}

	s.Seed(seed3...)
	for k, v := range wantSeed2Counts {
		assert.Equal(t, v, s.school[k])
	}
}

func TestSchool_CountFish(t *testing.T) {
	wantLen := 3
	input := []int{1, 2, 3}

	s := NewSchool(input)

	gotLen := s.CountFish()
	assert.Equal(t, wantLen, gotLen)
}

func TestSchool_Day(t *testing.T) {
	input := []int{3}

	d0Fish := map[int]int{3: 1}
	d0Count := 1

	d1Fish := map[int]int{2: 1}
	d1Count := 1

	d2Fish := map[int]int{1: 1}
	d2Count := 1

	d3Fish := map[int]int{0: 1}
	d3Count := 1

	d4Fish := map[int]int{6: 1, 8: 1}
	d4Count := 2

	s := NewSchool(input)
	assert.Equal(t, d0Count, s.CountFish())
	for k, got := range s.school {
		want, ok := d0Fish[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}

	s.Day()
	assert.Equal(t, d1Count, s.CountFish())
	for k, got := range s.school {
		want, ok := d1Fish[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}

	s.Day()
	assert.Equal(t, d2Count, s.CountFish())
	for k, got := range s.school {
		want, ok := d2Fish[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}

	s.Day()
	assert.Equal(t, d3Count, s.CountFish())
	for k, got := range s.school {
		want, ok := d3Fish[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}

	s.Day()
	assert.Equal(t, d4Count, s.CountFish())
	for k, got := range s.school {
		want, ok := d4Fish[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}
}

func TestSchool_ProcessDays(t *testing.T) {
	input := []int{3}
	wantLen := 2
	wantSchool := map[int]int{6: 1, 8: 1}

	s := NewSchool(input)

	s.ProcessDays(4)
	assert.Equal(t, wantLen, s.CountFish())
	for k, got := range s.school {
		want, ok := wantSchool[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}
}

func TestSchool_ProcessDays2(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	wantLen := 5
	wantSchool := map[int]int{0: 1, 1: 1, 2: 2, 3: 1}

	s := NewSchool(input)

	s.ProcessDays(1)
	assert.Equal(t, wantLen, s.CountFish())
	for k, got := range s.school {
		want, ok := wantSchool[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}
}

func TestSchool_ProcessDays3(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	wantLen := 26
	wantSchool := map[int]int{0: 3, 1: 5, 2: 3, 3: 2, 4: 2, 5: 1, 6: 5, 7: 1, 8: 4}

	s := NewSchool(input)

	s.ProcessDays(18)
	assert.Equal(t, wantLen, s.CountFish())
	for k, got := range s.school {
		want, ok := wantSchool[k]
		if !ok {
			assert.Equal(t, 0, got)
		} else {
			assert.Equal(t, want, got)
		}
	}
}

func TestSchool_ProcessDays4(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	wantLen := 5934

	s := NewSchool(input)

	s.ProcessDays(80)
	assert.Equal(t, wantLen, s.CountFish())
}

func TestSchool_ProcessDays5(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	wantLen := 26984457539

	s := NewSchool(input)

	s.ProcessDays(256)
	assert.Equal(t, wantLen, s.CountFish())
}
