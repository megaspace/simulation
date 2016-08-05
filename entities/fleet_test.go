package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/server/core"
	"github.com/megaspace/server/entities"
)

func TestAttachShip(t *testing.T) {
	ship := entities.Ship{10, 46.1}
	fleet := entities.NewFleet(1, "Test Fleet")
	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0, cap(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)

	attachedFleet := fleet.AttachShip(ship)
	// Arguments unaffected
	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)
	// Output has new value
	assert.Equal(t, 1, len(attachedFleet.Ships))
	assert.Equal(t, 46.1, attachedFleet.Speed)
}

func TestDetachShip(t *testing.T) {
	var ship entities.Ship
	var fleet, detachedFleet entities.Fleet

	ship = entities.Ship{10, 46.1}
	fleet = entities.Fleet{
		1,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]entities.Ship{ship},
		46.1,
	}
	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)

	detachedFleet = fleet.DetachShip(ship)
	// Arguments unaffected
	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)
	// Output has new value
	assert.Equal(t, 0, len(detachedFleet.Ships))
	assert.Equal(t, 0.0, detachedFleet.Speed)
}

func TestMoveTowards(t *testing.T) {
	var ship entities.Ship
	var fleet, movedFleet entities.Fleet

	ship = entities.Ship{345, 10.0}
	fleet = entities.Fleet{
		5234,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]entities.Ship{ship},
		10.0,
	}

	// Move a distance we can reach
	movedFleet = fleet.MoveTowards(core.Vector{1, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{0, 0, 0}, fleet.Coordinates)
	assert.Equal(t, core.Vector{1, 0, 0}, movedFleet.Coordinates)

	// Move to exactly the distance we can reach
	movedFleet = fleet.MoveTowards(core.Vector{10, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{0, 0, 0}, fleet.Coordinates)
	assert.Equal(t, core.Vector{10, 0, 0}, movedFleet.Coordinates)

	// Move a distance we can't reach
	movedFleet = fleet.MoveTowards(core.Vector{20, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{0, 0, 0}, fleet.Coordinates)
	assert.Equal(t, core.Vector{10, 0, 0}, movedFleet.Coordinates)
}
