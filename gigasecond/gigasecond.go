// Package gigasecond implements basic routines for
// working with one billion seconds.
package gigasecond

import "time"

// AddGigasecond returns a moment from t
// after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
