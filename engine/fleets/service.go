package fleets

import (
	"errors"
	"time"

	"github.com/megaspace/simulation/engine/ships"
)

// Service is the service layer for all things fleet related
type Service struct {
	nextFleetID FleetID
	fleets      map[FleetID]*Fleet
}

// NewService returns a new service to manage all the fleets.
// Only one of these is needed per game.
func NewService() *Service {
	fleetService := new(Service)
	fleetService.fleets = make(map[FleetID]*Fleet)
	return fleetService
}

func (fs *Service) CreateFleet(name string, initialShip *ships.Ship) FleetID {
	fleetID := fs.nextFleetID
	fs.nextFleetID = fs.nextFleetID + 1

	fleet := newFleet(fleetID, name)
	fleet.attachShip(initialShip)
	fs.fleets[fleetID] = fleet

	return fleetID
}

func (fs *Service) IssueOrder(fleetID FleetID, order Order) error {
	var ok bool
	var fleet *Fleet
	if fleet, ok = fs.fleets[fleetID]; !ok {
		return errors.New("unknown fleet")
	}
	fleet.issueOrder(order)
	return nil
}

func (fs *Service) Update(delta time.Duration) {
	for _, fleet := range fs.fleets {
		fleet.Update(delta)
	}
}
