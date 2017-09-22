package fleets

// OrderStatus shows whether the order is pending, or in progress, or complete, etc...
type OrderStatus byte

const (
	// OrderStatusNotAssigned means that the order has been created but not assigned to a Fleet
	OrderStatusNotAssigned OrderStatus = iota
	// OrderStatusPending means that the order is assigned to a Fleet but not yet started
	OrderStatusPending
	// OrderStatusInProgress means that the order is currently being executed
	OrderStatusInProgress
	// OrderStatusComplete means that the order is now complete
	OrderStatusComplete
)
