package fleets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/ships"
)

func TestNewMoveOrder_NOT_ASSIGNED(t *testing.T) {
	target := core.Vector{24, -23, 23.4}
	var order Order
	order = NewMoveOrder(target)

	assert.Equal(t, ORDER_NOT_ASSIGNED, order.Status())
}

func TestNewMoveOrder_PENDING(t *testing.T) {
	fleet := New(1, "Test Fleet")
	target := core.Vector{24, -23, 23.4}
	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	assert.Equal(t, ORDER_PENDING, order.Status())
}

func TestExecute_IN_PROGRESS(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{20, 0, 0}

	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)
	assert.Equal(t, ORDER_IN_PROGRESS, order.Status())
}

func TestExecute_COMPLETE(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{10, 0, 0}

	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, target, fleet.Coordinates)
	assert.Equal(t, ORDER_COMPLETE, order.Status())
}
