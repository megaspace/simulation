package fleets

import (
	"math"
	"time"

	"github.com/megaspace/simulation/core"
	"github.com/megaspace/simulation/ships"
)

type FleetId int64

type Fleet struct {
	Id          FleetId
	Name        string
	Coordinates core.Vector
	Ships       []*ships.Ship
	Speed       float64
}

func New(id FleetId, name string) *Fleet {
	fleet := new(Fleet)
	fleet.Id = id
	fleet.Name = name
	fleet.Coordinates = core.Vector{0, 0, 0}
	fleet.Ships = []*ships.Ship{}
	fleet.Speed = 0.0
	return fleet
}

func findShipIndex(s []*ships.Ship, id ships.ShipId) int {
	for i, ship := range s {
		if ship.Id == id {
			return i
		}
	}
	return -1
}

func (f *Fleet) updateSpeed() {
	if len(f.Ships) == 0 {
		f.Speed = 0.0
		return
	}

	lowestSpeed := math.MaxFloat64
	for _, ship := range f.Ships {
		if ship.Speed < lowestSpeed {
			lowestSpeed = ship.Speed
		}
	}

	f.Speed = lowestSpeed
}

func (f *Fleet) AttachShip(ship *ships.Ship) {
	f.Ships = append(f.Ships, ship)
	f.updateSpeed()
	return
}

func (f *Fleet) DetachShip(ship *ships.Ship) {
	i := findShipIndex(f.Ships, ship.Id)

	if i != -1 {
		f.Ships = append(f.Ships[:i], f.Ships[i+1:]...)
		f.updateSpeed()
	}

	return
}

func (f *Fleet) MoveTowards(target core.Vector, duration time.Duration) {
	distanceMoved := f.Speed * duration.Seconds()
	distanceVector := target.Subtract(f.Coordinates)
	movementVector := distanceVector.Normalize().Multiply(distanceMoved)

	if movementVector.Magnitude() > distanceVector.Magnitude() {
		f.Coordinates = target
	} else {
		f.Coordinates = f.Coordinates.Add(movementVector)
	}

	return
}
