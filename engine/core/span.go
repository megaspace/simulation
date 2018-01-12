package core

import (
	"math/rand"
)

type Span struct {
	lowerBound float64
	upperBound float64
}

func NewSpan(lowerBound, upperBound float64) Span {
	return Span{lowerBound: lowerBound, upperBound: upperBound}
}

func (s Span) Range() float64 {
	return s.upperBound - s.lowerBound
}

func (s Span) RandomValue() float64 {
	return rand.Float64()*s.Range() + s.lowerBound
}
