package core

import "time"

type Clock struct {
	lastUpdate time.Time
}

func (c *Clock) Start() {
	c.lastUpdate = time.Now()
}

func (c *Clock) Delta() (delta time.Duration) {
	now := time.Now()
	delta = now.Sub(c.lastUpdate)
	c.lastUpdate = now
	return
}
