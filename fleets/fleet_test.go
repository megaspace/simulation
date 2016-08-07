package fleets_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/fleets"
	"github.com/megaspace/simulation/ships"
)

func TestAttachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := fleets.New(1, "Test Fleet")
	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0, cap(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)

	fleet.AttachShip(ship)

	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)
}

func TestDetachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := fleets.Fleet{
		1,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]*ships.Ship{ship},
		46.1,
	}
	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)

	fleet.DetachShip(ship)

	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)
}

func TestMoveTowards(t *testing.T) {
	ship := ships.New(345, 10.0)
	fleet := fleets.Fleet{
		5234,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]*ships.Ship{ship},
		10.0,
	}

	// Move a distance we can reach
	fleet.MoveTowards(core.Vector{1, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{1, 0, 0}, fleet.Coordinates)

	fleet.Coordinates = core.Vector{0, 0, 0}
	// Move to exactly the distance we can reach
	fleet.MoveTowards(core.Vector{10, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)

	fleet.Coordinates = core.Vector{0, 0, 0}
	// Move a distance we can't reach
	fleet.MoveTowards(core.Vector{20, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)
}
