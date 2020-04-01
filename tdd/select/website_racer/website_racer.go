package race

import (
	"fmt"
	"net/http"
	"time"
)

func pingEndpoint(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

// Race ...
func Race(a string, b string) (string, error) {
	select {
	case <-pingEndpoint(a):
		return a, nil
	case <-pingEndpoint(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
