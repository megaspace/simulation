package fleets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/ships"
)

func TestAttachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := New(1, "Test Fleet")
	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0, cap(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)

	fleet.AttachShip(ship)

	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)
}

func TestDetachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := Fleet{
		1,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]*ships.Ship{ship},
		46.1,
		[]Order{},
	}
	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)

	fleet.DetachShip(ship)

	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)
}

func TestMoveTowards(t *testing.T) {
	ship := ships.New(345, 10.0)
	fleet := Fleet{
		5234,
		"Test Fleet",
		core.Vector{0, 0, 0},
		[]*ships.Ship{ship},
		10.0,
		[]Order{},
	}

	// Move a distance we can reach
	fleet.moveTowards(core.Vector{1, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{1, 0, 0}, fleet.Coordinates)

	fleet.Coordinates = core.Vector{0, 0, 0}
	// Move to exactly the distance we can reach
	fleet.moveTowards(core.Vector{10, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)

	fleet.Coordinates = core.Vector{0, 0, 0}
	// Move a distance we can't reach
	fleet.moveTowards(core.Vector{20, 0, 0}, time.Second)
	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)
}

func TestIssueOrder(t *testing.T) {
	fleet := New(1, "Test Fleet")
	order1 := NewMoveOrder(core.Vector{10, 10, 10})
	order2 := NewMoveOrder(core.Vector{20, 20, 20})

	assert.Equal(t, 0, len(fleet.orders))

	fleet.IssueOrder(order1)
	assert.Equal(t, 1, len(fleet.orders))
	assert.Equal(t, order1, fleet.orders[0])

	fleet.IssueOrder(order2)
	assert.Equal(t, order2, fleet.orders[1])
	assert.Equal(t, 2, len(fleet.orders))
}
