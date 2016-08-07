package fleets_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/fleets"
	"github.com/megaspace/simulation/ships"
)

func TestNewMoveOrder_PENDING(t *testing.T) {
	fleet := fleets.New(1, "Test Fleet")
	target := core.Vector{24, -23, 23.4}
	var order fleets.Order
	order = fleets.NewMoveOrder(fleet, target)

	assert.Equal(t, fleets.PENDING, order.Status())
}

func TestExecute_IN_PROGRESS(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := fleets.New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{20, 0, 0}

	var order fleets.Order
	order = fleets.NewMoveOrder(fleet, target)

	order.Execute(time.Second)

	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)
	assert.Equal(t, fleets.IN_PROGRESS, order.Status())
}

func TestExecute_COMPLETE(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := fleets.New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{10, 0, 0}

	var order fleets.Order
	order = fleets.NewMoveOrder(fleet, target)

	order.Execute(time.Second)

	assert.Equal(t, target, fleet.Coordinates)
	assert.Equal(t, fleets.COMPLETE, order.Status())
}
