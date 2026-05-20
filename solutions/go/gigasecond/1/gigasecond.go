package gigasecond

import (
    "time"
)

func AddGigasecond(t time.Time) time.Time {
	return time.Unix(t.Unix() + 1_000_000_000, 0)
}
