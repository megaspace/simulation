package orders

import "time"

type OrderStatus byte

const (
	PENDING OrderStatus = iota
	IN_PROGRESS
	COMPLETE
)

type Order interface {
	Status() OrderStatus
	Execute(duration time.Duration)
}
