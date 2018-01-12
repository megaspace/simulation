package fleets

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/megaspace/simulation/engine/core"
	"github.com/megaspace/simulation/engine/ships"
)

func TestAttachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := newFleet(1, "Test Fleet")
	assert.Equal(t, 0, len(fleet.Ships))
	assert.Equal(t, 0, cap(fleet.Ships))
	assert.Equal(t, 0.0, fleet.Speed)

	fleet.attachShip(ship)

	assert.Equal(t, 1, len(fleet.Ships))
	assert.Equal(t, 46.1, fleet.Speed)
}

func TestDetachShip(t *testing.T) {
	ship := ships.New(10, 46.1)
	fleet := Fleet{
		1,
		"Test Fleet",
		core.VectorZero,
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
		core.VectorZero,
		[]*ships.Ship{ship},
		10.0,
		[]Order{},
	}

	// Move a distance we can reach
	fleet.moveTowards(core.NewVector(1, 0, 0), time.Second)
	assert.Equal(t, core.NewVector(1, 0, 0), fleet.Coordinates)

	fleet.Coordinates = core.VectorZero
	// Move to exactly the distance we can reach
	fleet.moveTowards(core.NewVector(10, 0, 0), time.Second)
	assert.Equal(t, core.NewVector(10, 0, 0), fleet.Coordinates)

	fleet.Coordinates = core.VectorZero
	// Move a distance we can't reach
	fleet.moveTowards(core.NewVector(20, 0, 0), time.Second)
	assert.Equal(t, core.NewVector(10, 0, 0), fleet.Coordinates)
}

func TestIssueOrder(t *testing.T) {
	fleet := newFleet(1, "Test Fleet")
	order1 := NewMoveOrder(core.NewVector(10, 10, 10))
	order2 := NewMoveOrder(core.NewVector(20, 20, 20))

	assert.Equal(t, 0, len(fleet.orders))

	fleet.issueOrder(order1)
	assert.Equal(t, 1, len(fleet.orders))
	assert.Equal(t, order1, fleet.orders[0])

	fleet.issueOrder(order2)
	assert.Equal(t, order2, fleet.orders[1])
	assert.Equal(t, 2, len(fleet.orders))
}

func TestUpdateMultipleOrders(t *testing.T) {
	fleet := newFleet(1, "Test Fleet")
	ship := ships.New(1, 10)
	destination1 := core.NewVector(10, 0, 0)
	destination2 := core.NewVector(20, 0, 0)
	order1 := NewMoveOrder(destination1)
	order2 := NewMoveOrder(destination2)

	fleet.attachShip(ship)
	fleet.issueOrder(order1)
	fleet.issueOrder(order2)

	fleet.Update(time.Second * 10)
	assert.Equal(t, destination1, fleet.Coordinates)

	fleet.Update(time.Second * 10)
	assert.Equal(t, destination2, fleet.Coordinates)
}
