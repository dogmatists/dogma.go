// This is free and unencumbered software released into the public domain.

package dogma

import "math"
import "testing"
import "github.com/stretchr/testify/assert"

func TestAngleFromDegrees(t *testing.T) {
	assert.Equal(t, AngleFromDegrees(0).Degrees(), 0.0)
	assert.Equal(t, AngleFromDegrees(90).Degrees(), 90.0)
	assert.Equal(t, AngleFromDegrees(180).Degrees(), 180.0)
	assert.Equal(t, AngleFromDegrees(360).Degrees(), 360.0)
}

func TestAngleFromRadians(t *testing.T) {
	assert.Equal(t, AngleFromRadians(0).Radians(), 0.0)
	assert.Equal(t, AngleFromRadians(0.5*math.Pi).Radians(), 0.5*math.Pi)
	assert.Equal(t, AngleFromRadians(1*math.Pi).Radians(), 1*math.Pi)
	assert.Equal(t, AngleFromRadians(2*math.Pi).Radians(), 2*math.Pi)
}

func TestAngleFromTurns(t *testing.T) {
	assert.Equal(t, AngleFromTurns(0).Turns(), 0.0)
	assert.Equal(t, AngleFromTurns(0.25).Turns(), 0.25)
	assert.Equal(t, AngleFromTurns(0.5).Turns(), 0.5)
	assert.Equal(t, AngleFromTurns(1).Turns(), 1.0)
}
