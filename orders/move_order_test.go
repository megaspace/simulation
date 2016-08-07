package orders_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/entities"
	"github.com/megaspace/simulation/orders"
)

func TestNewMoveOrder_PENDING(t *testing.T) {
	fleet := entities.NewFleet(1, "Test Fleet")
	target := core.Vector{24, -23, 23.4}
	var order orders.Order
	order = orders.NewMoveOrder(fleet, target)

	assert.Equal(t, orders.PENDING, order.Status())
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
	assert.Equal(t, orders.IN_PROGRESS, order.Status())
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
	assert.Equal(t, orders.COMPLETE, order.Status())
}
