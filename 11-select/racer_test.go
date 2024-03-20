package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, returning faster one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != fastURL {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("returns error if server does'nt respond in with 2s", func(t *testing.T) {
		delayedServer := makeDelayServer(4 * time.Second)

		defer delayedServer.Close()
		_, err := ConfigurableRacer(delayedServer.URL, delayedServer.URL, 2*time.Second)
		if err == nil {
			t.Error("expected an error but did not get one")
		}

	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
