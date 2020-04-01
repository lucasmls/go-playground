package race

import (
	"net/http"
	"time"
)

func measureRequestTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

// Race ...
func Race(a string, b string) string {
	aDuration := measureRequestTime(a)
	bDuration := measureRequestTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}
