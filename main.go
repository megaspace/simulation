package main

import (
	"fmt"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/fleets"
	"github.com/megaspace/simulation/ships"
)

func main() {
	fleet1 := fleets.New(1, "First Fleet")
	fleet1.AttachShip(ships.New(1, 1))

	fleet1Destination := core.Vector{X: 10, Y: 0, Z: 0}
	fleet1MoveOrder := fleets.NewMoveOrder(fleet1Destination)
	fleet1.IssueOrder(fleet1MoveOrder)

	fleet1Destination2 := core.Vector{X: -5, Y: 10, Z: 1}
	fleet1MoveOrder2 := fleets.NewMoveOrder(fleet1Destination2)
	fleet1.IssueOrder(fleet1MoveOrder2)

	clock := core.Clock{}
	clock.Start()

	for {
		delta := clock.Delta()
		fleet1.Update(delta)
		fmt.Println(fleet1.Coordinates)
	}
}
