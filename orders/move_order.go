package orders

import (
	"time"

	"github.com/megaspace/server/core"
	"github.com/megaspace/server/entities"
)

type MoveOrder struct {
	fleet  *entities.Fleet
	target core.Vector
	Status OrderStatus
}

func NewMoveOrder(fleet *entities.Fleet, target core.Vector) *MoveOrder {
	order := new(MoveOrder)
	order.fleet = fleet
	order.target = target
	order.Status = PENDING
	return order
}

func (o *MoveOrder) Execute(duration time.Duration) {
	if o.Status == PENDING {
		o.Status = IN_PROGRESS
	}

	o.fleet.MoveTowards(o.target, duration)

	if o.fleet.Coordinates == o.target {
		o.Status = COMPLETE
	}
}

func (o *MoveOrder) IsPending() bool {
	return o.Status == PENDING
}

func (o *MoveOrder) IsInProgress() bool {
	return o.Status == IN_PROGRESS
}

func (o *MoveOrder) IsComplete() bool {
	return o.Status == COMPLETE
}
