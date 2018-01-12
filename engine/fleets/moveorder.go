package fleets

import (
	"time"

	"github.com/megaspace/simulation/engine/core"
)

// MoveOrder is a simple command to a Fleet to move to some position in space
type MoveOrder struct {
	fleet  *Fleet
	target core.Vector
	status OrderStatus
}

// NewMoveOrder is a constructor that just takes the target Vector
func NewMoveOrder(target core.Vector) Order {
	order := new(MoveOrder)
	order.target = target
	order.status = OrderStatusNotAssigned
	return order
}

// Status returns the current status of the order
func (o *MoveOrder) Status() OrderStatus {
	return o.status
}

func (o *MoveOrder) assignFleet(f *Fleet) {
	o.fleet = f
	o.status = OrderStatusPending
}

func (o *MoveOrder) execute(duration time.Duration) {
	if o.fleet == nil {
		return
	}

	if o.status == OrderStatusPending {
		o.status = OrderStatusInProgress
	}

	o.fleet.moveTowards(o.target, duration)

	if o.fleet.Coordinates == o.target {
		o.status = OrderStatusComplete
	}
}
