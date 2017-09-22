package fleets

type OrderStatus byte

const (
	OrderNotAssigned OrderStatus = iota
	OrderPending
	OrderInProgress
	OrderComplete
)
