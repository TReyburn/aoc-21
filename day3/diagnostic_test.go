package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunPowerDiagnostic(t *testing.T) {
	wantGamma := "10110"
	wantEpsilon := "01001"

	gotGamma, gotEpsilon := RunPowerDiagnostic(testData)

	assert.Equal(t, wantGamma, gotGamma)
	assert.Equal(t, wantEpsilon, gotEpsilon)
}

func TestFindLifeSupportRatings(t *testing.T) {
	wantO2Rating := "10111"
	wantCO2rating := "01010"

	gotO2Rating, gotCO2Rating := FindLifeSupportRatings(testData)

	assert.Equal(t, wantO2Rating, gotO2Rating)
	assert.Equal(t, wantCO2rating, gotCO2Rating)
}

func TestCalculatePower(t *testing.T) {
	wantPower := int64(198)
	testGamma := "10110"
	testEpsilon := "01001"

	gotPower := CalculatePower(testGamma, testEpsilon)

	assert.Equal(t, wantPower, gotPower)
}
