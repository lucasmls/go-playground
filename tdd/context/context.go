package context1

import (
	"fmt"
	"net/http"
)

// Store fetches data
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns a handler for calling Store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		dataCh := make(chan string, 1)

		go func() {
			dataCh <- store.Fetch()
		}()

		select {
		case d := <-dataCh:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
