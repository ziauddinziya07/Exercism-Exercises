package clock

import "fmt"

// Clock is struct type which demonstrates a clock with hours and minutes ....
type Clock struct {

	h, m int
}
var c Clock

// New returns the Clock which shows the given time ....
func New(h, m int) Clock {

	h += m / 60
	m = m % 60
	h %= 24
	if m < 0 {

		h --
		m += 60
	}
	if h < 0 {

		h += 24
	}
	c.h, c.m = h, m
	return c 
}

func (Clock) String() string {

	return fmt.Sprintf("%.2d:%.2d", c.h, c.m)
} 

// Add method adds the given minutes to the Clock return the Clock with updated time
func (Clock) Add(m int) Clock {

	return New(c.h, c.m + m)
}

// Subtract method subtracts the given minutes from the Clock return the Clock with updated time
func (Clock) Subtract(m int) Clock {

	return New(c.h, c.m - m)
}