package fleets

type OrderStatus byte

const (
	ORDER_NOT_ASSIGNED OrderStatus = iota
	ORDER_PENDING
	ORDER_IN_PROGRESS
	ORDER_COMPLETE
)
