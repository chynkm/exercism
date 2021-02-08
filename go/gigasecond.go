// Package gigasecond contains methods for calculating gigasecond
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = time.Second * 1000000000

// AddGigasecond will add a gigasecond to the provided time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(gigasecond)
	return t
}
