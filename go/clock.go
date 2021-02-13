package clock

import (
	"strconv"
)

// Clock structure
type Clock struct {
	h int
	m int
}

// New creates a clock time
func New(h, m int) *Clock {
	c := Clock{h, m}
	return &c
}

// String returns the string format of Clock in "hh:mm" format
func (c *Clock) String() string {
	var hour, min string
	// using a new variable for hour
	var extraH int

	// using a new variable for minute
	tempM := c.m
	if tempM < 0 {
		extraH = -1
		if tempM < -60 {
			extraH += tempM / 60
			tempM = tempM % 60
		}

		tempM = 60 + tempM
	}

	if tempM >= 60 {
		extraH = tempM / 60
		tempM = tempM % 60
	}

	tempH := c.h + extraH
	if tempH < 0 {
		if tempH <= -24 {
			tempH = tempH % 24
		}

		tempH = 24 + tempH
	}

	if tempH >= 24 {
		tempH = tempH % 24
	}

	if tempH <= 9 {
		hour = "0" + strconv.Itoa(tempH)
	} else {
		hour = strconv.Itoa(tempH)
	}

	if tempM <= 9 {
		min = "0" + strconv.Itoa(tempM)
	} else {
		min = strconv.Itoa(tempM)
	}

	return hour + ":" + min
}

// Add minutes to a Clock
func (c *Clock) Add(m int) *Clock {
	c.m += m
	if c.m >= 60 {
		c.h += c.m / 60
		c.m = c.m % 60
	}

	if c.h >= 24 {
		c.h = c.h % 24
	}

	return c
}
