package fleets

import "time"

type Order interface {
	Status() OrderStatus
	Execute(duration time.Duration)
}
