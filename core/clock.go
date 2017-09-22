package core

import "time"

// Clock is responsible for keeping track of game time between updates
type Clock struct {
	lastUpdate time.Time
}

// Start initializes the clock so that the delta will never get the duration from 0
func (c *Clock) Start() {
	c.lastUpdate = time.Now()
}

// Delta returns the duration of time since the last time Delta was invoked
func (c *Clock) Delta() (delta time.Duration) {
	now := time.Now()
	delta = now.Sub(c.lastUpdate)
	c.lastUpdate = now
	return
}
