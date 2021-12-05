package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriveSub(t *testing.T) {
	wantHPos := 15
	wantDepth := 10

	gotHPos, gotDepth := DriveSub(testData)

	assert.Equal(t, wantHPos, gotHPos)
	assert.Equal(t, wantDepth, gotDepth)
}

func TestDriveSubWithAim(t *testing.T) {
	wantHPos := 15
	wantDepth := 60

	gotHPos, gotDepth := DriveSubWithAim(testData)

	assert.Equal(t, wantHPos, gotHPos)
	assert.Equal(t, wantDepth, gotDepth)
}

func TestDriveOutcome(t *testing.T) {
	wantOutcome := 150

	gotOutcome := DriveOutcome(DriveSub(testData))

	assert.Equal(t, wantOutcome, gotOutcome)
}
