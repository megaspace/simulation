package fleets

import (
	"testing"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/ships"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreateFleet(t *testing.T) {
	fleetService := NewFleetService()

	ship := ships.New(1, 1)
	fleetID := fleetService.CreateFleet("test", ship)
	assert.Equal(t, FleetID(0), fleetID)

	fleet, _ := fleetService.fleets[fleetID]
	assert.Equal(t, "test", fleet.Name)
	assert.Equal(t, []*ships.Ship{ship}, fleet.Ships)

	ship2 := ships.New(1, 1)
	fleetID2 := fleetService.CreateFleet("test2", ship2)
	assert.Equal(t, FleetID(1), fleetID2)

	fleet2, _ := fleetService.fleets[fleetID2]
	assert.Equal(t, "test2", fleet2.Name)
	assert.Equal(t, []*ships.Ship{ship2}, fleet2.Ships)
}

func TestServiceIssueOrder(t *testing.T) {
	fleetService := NewFleetService()

	fleet := newFleet(0, "test")
	fleetService.fleets = map[FleetID]*Fleet{fleet.ID: fleet}

	order := NewMoveOrder(core.VectorZero)
	fleetService.IssueOrder(0, order)

	assert.Equal(t, []Order{order}, fleet.orders)
	assert.Equal(t, OrderStatusPending, fleet.orders[0].Status())
}