package stars

import (
	"github.com/megaspace/simulation/engine/core"
)

// StarID is a unique int64 identifying the Star
type StarID int64

// Star is a combusting ball of gas!
type Star struct {
	ID             StarID
	Name           string
	Coordinates    core.Vector
	Classification string
	Temperature    float64
	Mass           float64
	Radius         float64
	Color          core.Color
}
