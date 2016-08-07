package orders

import (
	"time"

	"github.com/megaspace/server/core"
	"github.com/megaspace/server/entities"
)

type MoveOrder struct {
	fleet  *entities.Fleet
	target core.Vector
	status OrderStatus
}

func NewMoveOrder(fleet *entities.Fleet, target core.Vector) Order {
	order := new(MoveOrder)
	order.fleet = fleet
	order.target = target
	order.status = PENDING
	return order
}

func (o *MoveOrder) Status() OrderStatus {
	return o.status
}

func (o *MoveOrder) Execute(duration time.Duration) {
	if o.status == PENDING {
		o.status = IN_PROGRESS
	}

	o.fleet.MoveTowards(o.target, duration)

	if o.fleet.Coordinates == o.target {
		o.status = COMPLETE
	}
}