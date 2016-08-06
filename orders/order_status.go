package orders

type OrderStatus byte

const (
	PENDING OrderStatus = iota
	IN_PROGRESS
	COMPLETE
)
