package clock

import (
    "fmt"
    "math"
)

type Clock struct {
    hours, minutes int
}

func New(h, m int) Clock {
	clock := Clock {hours: h, minutes: m}

    absoluteMinutes := clock.getAbsoluteMinutes()
	absoluteHours := clock.getAbsoluteHours(absoluteMinutes)
    
	normalizedHours := (absoluteHours % 24 + 24) % 24
    normalizedMinutes := (absoluteMinutes % 60 + 60) % 60
    
	return Clock {
        hours: normalizedHours, 
        minutes: normalizedMinutes,
    }
}

func (c Clock) Add(m int) Clock {
	absoluteMinutes := c.getAbsoluteMinutes() + m
	absoluteHours := c.getAbsoluteHours(absoluteMinutes)
    
	return Clock {
        hours: (absoluteHours % 24 + 24) % 24,
        minutes: (absoluteMinutes % 60 + 60) % 60,
    }
}

func (c Clock) Subtract(m int) Clock {
	absoluteMinutes := c.getAbsoluteMinutes() - m
	absoluteHours := c.getAbsoluteHours(absoluteMinutes)
    
	return Clock {
        hours: (absoluteHours % 24 + 24) % 24,
        minutes: (absoluteMinutes % 60 + 60) % 60,
    }
}

func (c Clock) String() string {    
    absoluteMinutes := c.getAbsoluteMinutes()
	absoluteHours := c.getAbsoluteHours(absoluteMinutes)
    
	normalizedHours := (absoluteHours % 24 + 24) % 24
    normalizedMinutes := (absoluteMinutes % 60 + 60) % 60
    
	return fmt.Sprintf("%02d:%02d", normalizedHours, normalizedMinutes)
}

func (c Clock) getAbsoluteMinutes() int {
    return c.hours * 60 + c.minutes
}

func (c Clock) getAbsoluteHours(absoluteMinutes int) int {
	absoluteHours := int(absoluteMinutes / 60)

    if absoluteMinutes < 0 {
        absoluteHours = -1 * int(math.Ceil(-1 * float64(absoluteMinutes) / 60))
    }

    return absoluteHours
}