package fleets

import "time"

// Order is an interface for every kind of order that can be issued to a Fleet.
type Order interface {
	Status() OrderStatus

	execute(duration time.Duration)
	assignFleet(f *Fleet)
}
