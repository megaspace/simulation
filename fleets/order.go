package fleets

import "time"

type Order interface {
	Status() OrderStatus

	execute(duration time.Duration)
	assignFleet(f *Fleet)
}
