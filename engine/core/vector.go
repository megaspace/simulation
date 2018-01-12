package core

import (
	"math"
	"math/rand"
)

var (
	// VectorZero is a vector with 0 magnitude
	VectorZero = Vector{X: 0, Y: 0, Z: 0}
)

// Vector is a 3-dimensional vector with an X, Y and Z component
type Vector struct {
	X, Y, Z float64
}

// NewVector is a helper function for constructing a new Vector
func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

// Add is a pure function addition of two vectors
func (v Vector) Add(operand Vector) Vector {
	v.X += operand.X
	v.Y += operand.Y
	v.Z += operand.Z
	return v
}

// Subtract is a pure function subtraction of two vectors
func (v Vector) Subtract(operand Vector) Vector {
	v.X -= operand.X
	v.Y -= operand.Y
	v.Z -= operand.Z
	return v
}

// Magnitude gives the length of the vector
func (v Vector) Magnitude() float64 {
	pow := math.Pow
	sqrt := math.Sqrt
	return sqrt(pow(v.X, 2) + pow(v.Y, 2) + pow(v.Z, 2))
}

// Divide reduces the length of the vector by the scalar
func (v Vector) Divide(scalar float64) Vector {
	v.X /= scalar
	v.Y /= scalar
	v.Z /= scalar
	return v
}

// Multiply increases the length of the vector by the scalar
func (v Vector) Multiply(scalar float64) Vector {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
	return v
}

// Normalize is a pure function that returns a unit vector in the same direction as the original
func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()

	if magnitude == 0 {
		return v
	}

	return v.Divide(magnitude)
}

func CreateVectorWithinRadius(radius float64) Vector {
	randomValueX := rand.Float64() - 0.5
	randomValueY := rand.Float64() - 0.5
	randomValueZ := rand.Float64() - 0.5

	randomValueRadial := rand.Float64()

	directionVector := NewVector(randomValueX, randomValueY, randomValueZ).Normalize()

	return directionVector.Multiply(randomValueRadial * radius)
}
