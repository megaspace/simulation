package core

import (
	"math"
)

var (
	VectorZero = Vector{X: 0, Y: 0, Z: 0}
)

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

func (v1 Vector) Add(v2 Vector) Vector {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
	return v1
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z
	return v1
}

func (v Vector) Magnitude() float64 {
	pow := math.Pow
	sqrt := math.Sqrt
	return sqrt(pow(v.X, 2) + pow(v.Y, 2) + pow(v.Z, 2))
}

func (v Vector) Divide(scalar float64) Vector {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
	return v
}

func (v Vector) Multiply(scalar float64) Vector {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	return v
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return v
	}

	return v.Divide(magnitude)
}
