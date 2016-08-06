package entities

type ShipId int64

type Ship struct {
	Id    ShipId
	Speed float64
}

func NewShip(id ShipId, speed float64) *Ship {
	ship := new(Ship)
	ship.Id = id
	ship.Speed = speed
	return ship
}
