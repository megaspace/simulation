package main

import (
	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/fleets"
	"github.com/megaspace/simulation/ships"
)

func main() {
	fleetService := fleets.NewFleetService()

	ship1 := ships.New(1, 0.00000165343)
	ship2 := ships.New(2, 0.004352)

	fleet1ID := fleetService.CreateFleet("Slow Fleet", ship1)
	fleet2ID := fleetService.CreateFleet("Fast Fleet", ship2)

	fleet1MoveOrder := fleets.NewMoveOrder(core.NewVector(1, 0, 0))
	if err := fleetService.IssueOrder(fleet1ID, fleet1MoveOrder); err != nil {
		panic(err)
	}

	fleet2MoveOrder := fleets.NewMoveOrder(core.NewVector(-5, 10, 1))
	if err := fleetService.IssueOrder(fleet2ID, fleet2MoveOrder); err != nil {
		panic(err)
	}

	clock := core.Clock{}
	clock.Start()

	for {
		delta := clock.Delta()
		fleetService.Update(delta)
	}
}
