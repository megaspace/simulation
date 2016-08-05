package core

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	v1.X = v1.X + v2.X
	v1.Y = v1.Y + v2.Y
	v1.Z = v1.Z + v2.Z
	return v1
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	v1.X = v1.X - v2.X
	v1.Y = v1.Y - v2.Y
	v1.Z = v1.Z - v2.Z
	return v1
}

func (v Vector) Magnitude() float64 {
	pow := math.Pow
	sqrt := math.Sqrt
	return sqrt(pow(v.X, 2) + pow(v.Y, 2) + pow(v.Z, 2))
}

func (v Vector) Divide(scalar float64) Vector {
	v.X = v.X / scalar
	v.Y = v.Y / scalar
	v.Z = v.Z / scalar
	return v
}

func (v Vector) Multiply(scalar float64) Vector {
	v.X = v.X * scalar
	v.Y = v.Y * scalar
	v.Z = v.Z * scalar
	return v
}

func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return v
	}

	return v.Divide(magnitude)
}
