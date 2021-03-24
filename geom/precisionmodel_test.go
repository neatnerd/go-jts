package geom

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestParameterlessConstructor(t *testing.T) {
	p := NewDefaultPrecisionModel()
	assert2.Equal(t, 0.0, p.scale)
}

func TestGetMaximumSignificantDigits(t *testing.T) {
	assert := assert2.New(t)
	floating := NewPrecisionModel(FLOATING)
	floating_single := NewPrecisionModel(FLOATING_SINGLE)
	fixed := NewPrecisionModel(FIXED)
	fixedN := NewFixedPrecisionModel(1000)
	assert.Equal(16, floating.getMaximumSignificantDigits())
	assert.Equal(6, floating_single.getMaximumSignificantDigits())
	assert.Equal(1, fixed.getMaximumSignificantDigits())
	assert.Equal(4, fixedN.getMaximumSignificantDigits())
}

func TestMakePrecise(t *testing.T) {
	pm_10 := NewFixedPrecisionModel(0.1)
	precisionCoordinateTester(pm_10, 1200.4, 1240.4, 1200, 1240, t)
	precisionCoordinateTester(pm_10, 1209.4, 1240.4, 1210, 1240, t)
}

func precisionCoordinateTester(pm PrecisionModel, x1, y1, x2, y2 float64, t *testing.T) {
	p := NewXYCoordinate(x1, y1)
	pm.makePreciseCoordinate(&p)
	pPrecise := NewXYCoordinate(x2, y2)
	assert2.True(t, p.equals2D(&pPrecise))
}
