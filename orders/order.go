package orders

import "time"

type OrderStatus byte

const (
	PENDING OrderStatus = iota
	IN_PROGRESS
	COMPLETE
)

type Order interface {
	Execute(duration time.Duration)

	IsPending() bool
	IsInProgress() bool
	IsComplete() bool
}
