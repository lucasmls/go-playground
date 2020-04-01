package race

import (
	"net/http"
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
func Race(a string, b string) string {
	select {
	case <-pingEndpoint(a):
		return a
	case <-pingEndpoint(b):
		return b
	}
}
