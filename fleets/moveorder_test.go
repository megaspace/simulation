package fleets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/ships"
)

func TestNewMoveOrderStatusNotAssigned(t *testing.T) {
	target := core.NewVector(24, -23, 23.4)
	var order Order
	order = NewMoveOrder(target)

	assert.Equal(t, OrderStatusNotAssigned, order.Status())
}

func TestNewMoveOrderStatusPending(t *testing.T) {
	fleet := New(1, "Test Fleet")
	target := core.NewVector(24, -23, 23.4)
	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	assert.Equal(t, OrderStatusPending, order.Status())
}

func TestExecute_IN_PROGRESS(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.NewVector(20, 0, 0)

	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, core.NewVector(10, 0, 0), fleet.Coordinates)
	assert.Equal(t, OrderStatusInProgress, order.Status())
}

func TestExecute_COMPLETE(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := New(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.NewVector(10, 0, 0)

	var order Order
	order = NewMoveOrder(target)
	fleet.IssueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, target, fleet.Coordinates)
	assert.Equal(t, OrderStatusComplete, order.Status())
}
