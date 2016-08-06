package orders_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/server/core"
	"github.com/megaspace/server/entities"
	"github.com/megaspace/server/orders"
)

func TestNewMoveOrder_PENDING(t *testing.T) {
	fleet := entities.NewFleet(1, "Test Fleet")
	target := core.Vector{24, -23, 23.4}
	var order orders.Order
	order = orders.NewMoveOrder(fleet, target)

	assert.True(t, order.IsPending())
}

func TestExecute_IN_PROGRESS(t *testing.T) {
	ship := entities.NewShip(1, 10)
	fleet := entities.NewFleet(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{20, 0, 0}

	var order orders.Order
	order = orders.NewMoveOrder(fleet, target)

	order.Execute(time.Second)

	assert.Equal(t, core.Vector{10, 0, 0}, fleet.Coordinates)
	assert.True(t, order.IsInProgress())
}

func TestExecute_COMPLETE(t *testing.T) {
	ship := entities.NewShip(1, 10)
	fleet := entities.NewFleet(1, "Test Fleet")
	fleet.AttachShip(ship)
	target := core.Vector{10, 0, 0}

	var order orders.Order
	order = orders.NewMoveOrder(fleet, target)

	order.Execute(time.Second)

	assert.Equal(t, target, fleet.Coordinates)
	assert.True(t, order.IsComplete())
}
