package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseData(t *testing.T) {
	data := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe`
	wantInputsLen := 1
	wantInputLen := 10
	wantOutputsLen := 1
	wantOutputLen := 4

	gotInputs, gotOutputs := ParseData(data)
	assert.Equal(t, wantInputsLen, len(gotInputs))
	assert.Equal(t, wantInputLen, len(gotInputs[0]))
	assert.Equal(t, wantOutputsLen, len(gotOutputs))
	assert.Equal(t, wantOutputLen, len(gotOutputs[0]))
}

func TestParseData_MultiRow(t *testing.T) {
	data := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
`
	wantInputsLen := 2
	wantInputLen := 10
	wantOutputsLen := 2
	wantOutputLen := 4

	gotInputs, gotOutputs := ParseData(data)
	assert.Equal(t, wantInputsLen, len(gotInputs))
	assert.Equal(t, wantInputLen, len(gotInputs[0]))
	assert.Equal(t, wantInputLen, len(gotInputs[1]))
	assert.Equal(t, wantOutputsLen, len(gotOutputs))
	assert.Equal(t, wantOutputLen, len(gotOutputs[0]))
	assert.Equal(t, wantOutputLen, len(gotOutputs[1]))
}
