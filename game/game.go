package game

import (
	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/fleets"
	"github.com/megaspace/simulation/ships"
	"github.com/megaspace/simulation/stars"
)

type Game struct {
	clock        *core.Clock
	fleetService *fleets.Service
	starService  *stars.Service
}

func New() (g *Game) {
	g = new(Game)
	g.clock = &core.Clock{}

	g.starService = stars.NewService()
	g.starService.GenerateGalaxy(200, 1000)

	g.fleetService = fleets.NewService()

	ship1 := ships.New(1, 0.00000165343)
	ship2 := ships.New(2, 0.004352)

	fleet1ID := g.fleetService.CreateFleet("Slow Fleet", ship1)
	fleet2ID := g.fleetService.CreateFleet("Fast Fleet", ship2)

	fleet1MoveOrder := fleets.NewMoveOrder(core.CreateVectorWithinRadius(1000))
	if err := g.fleetService.IssueOrder(fleet1ID, fleet1MoveOrder); err != nil {
		panic(err)
	}

	fleet2MoveOrder := fleets.NewMoveOrder(core.CreateVectorWithinRadius(1000))
	if err := g.fleetService.IssueOrder(fleet2ID, fleet2MoveOrder); err != nil {
		panic(err)
	}

	return
}

func (g *Game) Start() {
	g.clock.Start()

	for {
		delta := g.clock.Delta()
		g.fleetService.Update(delta)
	}
}
