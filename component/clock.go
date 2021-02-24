package component

import "time"

// Clock is a clock format
type Clock struct {
	Format string
}

// NewClock makes a clock.
// Format strings: Mon Jan 2 15:04:05 -0700 MST 2006
func NewClock(format string) *Clock {
	return &Clock{format}
}

// Title implements stat.Statiface
func (c *Clock) Title() string {
	return ""
}

// Check implements stat.Statiface
func (c *Clock) Check() string {
	return time.Now().Format(c.Format)
}
