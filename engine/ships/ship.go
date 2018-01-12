package ships

// ShipID is an int64 that is a unique reference to a ship
type ShipID int64

// Ship represents a literal spaceship
type Ship struct {
	ID    ShipID
	Speed float64
}

// New is a constructor for ship
func New(id ShipID, speed float64) *Ship {
	ship := new(Ship)
	ship.ID = id
	ship.Speed = speed
	return ship
}
