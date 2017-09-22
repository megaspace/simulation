package core

import (
	"time"
)

// Clock is responsible for keeping track of game time between updates
type Clock struct {
	lastUpdate  time.Time
	speedFactor int64
}

// Start initializes the clock so that the delta will never get the duration from 0
func (c *Clock) Start() {
	c.lastUpdate = time.Now()
	if c.speedFactor == 0 {
		c.speedFactor = 1
	}
}

// Delta returns the duration of time since the last time Delta was invoked
func (c *Clock) Delta() (delta time.Duration) {
	now := time.Now()
	delta = time.Duration(int64(now.Sub(c.lastUpdate)) * c.speedFactor)
	c.lastUpdate = now
	return
}

func (c *Clock) SetSpeedNormal() {
	c.speedFactor = 1
}

func (c *Clock) SetSpeedFast() {
	c.speedFactor = 10
}

func (c *Clock) SetSpeedVeryFast() {
	c.speedFactor = 100
}

func (c *Clock) SetSpeedSuperFast() {
	c.speedFactor = 1000
}
