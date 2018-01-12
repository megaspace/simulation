package fleets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/engine/core"
	"github.com/megaspace/simulation/engine/ships"
)

func TestNewMoveOrderStatusNotAssigned(t *testing.T) {
	target := core.NewVector(24, -23, 23.4)
	var order Order
	order = NewMoveOrder(target)

	assert.Equal(t, OrderStatusNotAssigned, order.Status())
}

func TestNewMoveOrderStatusPending(t *testing.T) {
	fleet := newFleet(1, "Test Fleet")
	target := core.NewVector(24, -23, 23.4)
	var order Order
	order = NewMoveOrder(target)
	fleet.issueOrder(order)

	assert.Equal(t, OrderStatusPending, order.Status())
}

func TestExecute_IN_PROGRESS(t *testing.T) {
	ship := ships.New(1, 5)
	fleet := newFleet(1, "Test Fleet")
	fleet.attachShip(ship)
	target := core.NewVector(100, 0, 0)

	var order Order
	order = NewMoveOrder(target)
	fleet.issueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, core.NewVector(5, 0, 0), fleet.Coordinates)
	assert.Equal(t, OrderStatusInProgress, order.Status())
}

func TestExecute_COMPLETE(t *testing.T) {
	ship := ships.New(1, 10)
	fleet := newFleet(1, "Test Fleet")
	fleet.attachShip(ship)
	target := core.NewVector(10, 0, 0)

	var order Order
	order = NewMoveOrder(target)
	fleet.issueOrder(order)

	order.execute(time.Second)

	assert.Equal(t, target, fleet.Coordinates)
	assert.Equal(t, OrderStatusComplete, order.Status())
}
