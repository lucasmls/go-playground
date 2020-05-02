package context1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	data := "Hello, world"

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {

		store := &SpyStore{
			response: data,
			t:        t,
		}

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})

	t.Run("returns data from the store", func(t *testing.T) {
		store := &SpyStore{
			response: data,
			t:        t,
		}

		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", wat "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})
}
