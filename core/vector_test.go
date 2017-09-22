package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
)

func TestAdd(t *testing.T) {
	v1 := core.NewVector(1, 2, 4)
	v2 := core.NewVector(8, 16, 32)
	v3 := v1.Add(v2)

	assert.Equal(t, core.NewVector(1, 2, 4), v1)
	assert.Equal(t, core.NewVector(8, 16, 32), v2)
	assert.Equal(t, core.NewVector(9, 18, 36), v3)
}

func TestSubtract(t *testing.T) {
	v1 := core.NewVector(4, 6, 2)
	v2 := core.NewVector(26, -53, 1)
	v3 := v1.Subtract(v2)

	assert.Equal(t, core.NewVector(4, 6, 2), v1)
	assert.Equal(t, core.NewVector(26, -53, 1), v2)
	assert.Equal(t, core.NewVector(-22, 59, 1), v3)
}

func TestMagnitude(t *testing.T) {
	var v core.Vector
	var mag float64

	v = core.VectorZero
	mag = v.Magnitude()
	assert.Equal(t, core.VectorZero, v)
	assert.Equal(t, 0.0, mag)

	v = core.NewVector(3, 4, 0)
	mag = v.Magnitude()
	assert.Equal(t, core.NewVector(3, 4, 0), v)
	assert.Equal(t, 5.0, mag)

	v = core.NewVector(-20, 0, 21)
	mag = v.Magnitude()
	assert.Equal(t, core.NewVector(-20, 0, 21), v)
	assert.Equal(t, 29.0, mag)
}

func TestDivide(t *testing.T) {
	var v core.Vector
	var result core.Vector

	v = core.NewVector(1, 5, -3)
	result = v.Divide(2)
	assert.Equal(t, core.NewVector(1, 5, -3), v)
	assert.Equal(t, core.NewVector(0.5, 2.5, -1.5), result)

	v = core.NewVector(5.42, -2524.4422, 0)
	result = v.Divide(4)
	assert.Equal(t, core.NewVector(5.42, -2524.4422, 0), v)
	assert.Equal(t, core.NewVector(1.355, -631.11055, 0), result)
}

func TestMultiply(t *testing.T) {
	var v core.Vector
	var result core.Vector

	v = core.NewVector(4.35, -3.223, 0)
	result = v.Multiply(2)
	assert.Equal(t, core.NewVector(4.35, -3.223, 0), v)
	assert.Equal(t, core.NewVector(8.7, -6.446, 0.0), result)

	v = core.NewVector(425, -23.2, 0.444)
	result = v.Multiply(0.1)
	assert.Equal(t, core.NewVector(425, -23.2, 0.444), v)
	assert.Equal(t, core.NewVector(42.5, -2.32, 0.0444), result)
}

func TestNormalize(t *testing.T) {
	var v, v2 core.Vector

	v = core.NewVector(3, 4, 0)
	v2 = v.Normalize()
	assert.Equal(t, core.NewVector(3, 4, 0), v)
	assert.Equal(t, 5.0, v.Magnitude())
	assert.Equal(t, 1.0, v2.Magnitude())

	v = core.VectorZero
	v2 = v.Normalize()
	assert.Equal(t, core.VectorZero, v)
	assert.Equal(t, 0.0, v.Magnitude())
	assert.Equal(t, 0.0, v2.Magnitude())
}
