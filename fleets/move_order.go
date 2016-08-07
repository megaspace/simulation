package fleets

import (
	"time"

	"github.com/megaspace/simulation/core"
)

type MoveOrder struct {
	fleet  *Fleet
	target core.Vector
	status OrderStatus
}

func NewMoveOrder(target core.Vector) Order {
	order := new(MoveOrder)
	order.target = target
	order.status = ORDER_NOT_ASSIGNED
	return order
}

func (o *MoveOrder) Status() OrderStatus {
	return o.status
}

func (o *MoveOrder) assignFleet(f *Fleet) {
	o.fleet = f
	o.status = ORDER_PENDING
}

func (o *MoveOrder) execute(duration time.Duration) {
	if o.fleet == nil {
		return
	}

	if o.status == ORDER_PENDING {
		o.status = ORDER_IN_PROGRESS
	}

	o.fleet.moveTowards(o.target, duration)

	if o.fleet.Coordinates == o.target {
		o.status = ORDER_COMPLETE
	}
}
