package day8

import (
	"github.com/TReyburn/aoc-21/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScreen_Process(t *testing.T) {
	data := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe`
	_, outputs := ParseData(data)

	s := NewScreen()

	for _, out := range outputs {
		s.Process(out)
	}

	want := map[int]int{4: 1, 8: 1}
	for i := 0; i < 10; i++ {
		wantCount, ok := want[i]
		if !ok {
			_, screenOk := s.nums[i]
			assert.False(t, screenOk)
		} else {
			screenCount, screenOk := s.nums[i]
			assert.True(t, screenOk)
			assert.Equal(t, wantCount, screenCount.count)
		}
	}
}

func TestScreen_Process2(t *testing.T) {
	data := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg`
	_, outputs := ParseData(data)

	s := NewScreen()

	for _, out := range outputs {
		s.Process(out)
	}

	want := map[int]int{1: 3, 4: 1, 7: 2, 8: 2}
	for i := 0; i < 10; i++ {
		wantCount, ok := want[i]
		if !ok {
			_, screenOk := s.nums[i]
			assert.False(t, screenOk)
		} else {
			screenCount, screenOk := s.nums[i]
			assert.True(t, screenOk)
			assert.Equal(t, wantCount, screenCount.count)
		}
	}
}

func TestScreen_CountNums(t *testing.T) {
	var testCase = []struct {
		want   int
		search []int
	}{
		{5, []int{1}},
		{0, []int{6, 7, 8}},
		{6, []int{1, 2, 3}},
	}

	for _, tc := range testCase {
		s := NewScreen()

		s.nums[1] = &Number{count: 5}
		s.nums[2] = &Number{count: 1}
		s.nums[3] = &Number{}

		got := s.CountNums(tc.search)
		assert.Equal(t, tc.want, got)
	}
}

func TestScreen_FullRun(t *testing.T) {
	want := 26
	search := []int{1, 4, 7, 8}

	s := NewScreen()

	_, outputs := ParseData(testData)

	for _, out := range outputs {
		s.Process(out)
	}

	got := s.CountNums(search)
	assert.Equal(t, want, got)
}

func TestScreen_LoadInstructions(t *testing.T) {
	var testCase = []struct {
		wantStr string
		wantInt int
	}{
		{algorithms.SortString("acedgfb"), 8},
		{algorithms.SortString("cdfbe"), 5},
		{algorithms.SortString("gcdfa"), 2},
		{algorithms.SortString("fbcad"), 3},
		{algorithms.SortString("dab"), 7},
		{algorithms.SortString("cefabd"), 9},
		{algorithms.SortString("cdfgeb"), 6},
		{algorithms.SortString("eafb"), 4},
		{algorithms.SortString("cagedb"), 0},
		{algorithms.SortString("ab"), 1},
	}
	data := `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`
	s := NewScreen()
	input, _ := ParseData(data)

	s.LoadInstructions(input[0])
	for _, tc := range testCase {
		assert.Equal(t, tc.wantInt, s.lookup[tc.wantStr])
		assert.Equal(t, tc.wantStr, s.reverseLookup[tc.wantInt])
	}

}

func TestScreen_ProcessFromInstructions(t *testing.T) {
	var testCase = []struct {
		want int
		data string
	}{
		{8394, `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe`},
		{9781, `edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc`},
		{1197, `fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg`},
		{9361, `fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb`},
		{4873, `aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea`},
		{8418, `fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb`},
		{4548, `dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe`},
		{1625, `bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef`},
		{8717, `egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb`},
		{4315, `gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`},
	}

	for _, tc := range testCase {
		s := NewScreen()

		input, output := ParseData(tc.data)

		s.LoadInstructions(input[0])

		got := s.ProcessFromInstructions(output[0])

		assert.Equal(t, tc.want, got)
	}
}

func TestScreen_FullRun2(t *testing.T) {
	want := 61229
	var got int

	inputs, outputs := ParseData(testData)

	for idx, out := range outputs {
		s := NewScreen()

		s.LoadInstructions(inputs[idx])
		got += s.ProcessFromInstructions(out)
	}

	assert.Equal(t, want, got)
}
