package fleets

import (
	"math"
	"time"

	"github.com/megaspace/simulation/engine/core"
	"github.com/megaspace/simulation/engine/ships"
)

// FleetID is a unique int64 identifying the Fleet
type FleetID int64

// Fleet is a collection of ships
// The speed of the whole Fleet is as fast the slowest ship included
type Fleet struct {
	ID          FleetID
	Name        string
	Coordinates core.Vector
	Ships       []*ships.Ship
	Speed       float64

	orders []Order
}

// New is a constructor for a Fleet that takes a FleetID (or int64) and a name
func newFleet(id FleetID, name string) *Fleet {
	fleet := new(Fleet)
	fleet.ID = id
	fleet.Name = name
	fleet.Coordinates = core.VectorZero
	fleet.Ships = []*ships.Ship{}
	fleet.Speed = 0.0
	return fleet
}

func (f *Fleet) attachShip(ship *ships.Ship) {
	f.Ships = append(f.Ships, ship)
	f.updateSpeed()
	return
}

// DetachShip removes the ship from the fleet list
func (f *Fleet) DetachShip(ship *ships.Ship) {
	i := findShipIndex(f.Ships, ship.ID)

	if i != -1 {
		f.Ships = append(f.Ships[:i], f.Ships[i+1:]...)
		f.updateSpeed()
	}

	return
}

func (f *Fleet) issueOrder(order Order) {
	f.orders = append(f.orders, order)
	order.assignFleet(f)
}

// Update is called as part of the main game loop to activate the fleet
func (f *Fleet) Update(duration time.Duration) {
	if len(f.orders) == 0 {
		return
	}

	for _, o := range f.orders {
		if o.Status() == OrderStatusComplete {
			continue
		}
		o.execute(duration)
		return
	}
}

func findShipIndex(s []*ships.Ship, id ships.ShipID) int {
	for i, ship := range s {
		if ship.ID == id {
			return i
		}
	}
	return -1
}

func (f *Fleet) moveTowards(target core.Vector, duration time.Duration) {
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
