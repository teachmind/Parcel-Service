package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	bindServer := NewServer(":1000", nil)
	go bindServer.Run()
	defer bindServer.Shutdown()
	time.Sleep(1 * time.Second)

	t.Run("test success run server", func(t *testing.T) {
		s := NewServer(":2000", nil)
		assert.NotNil(t, s)

		go func() {
			if err := s.Run(); err != nil {
				t.Error(err)
			}
		}()

		assert.NoError(t, s.Shutdown())
	})

	t.Run("test failed run with gracefully shutdown", func(t *testing.T) {
		s := NewServer(":1000", nil)
		assert.NotNil(t, s)
		assert.NoError(t, s.Run())
		assert.NoError(t, s.Shutdown())
	})
}

func TestPingHandler(t *testing.T) {
	t.Run("test server ping handler", func(t *testing.T) {
		s := NewServer(":2000", nil).route()
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/ping", nil)

		s.ServeHTTP(w, r)

		assert.Equal(t, w.Body.String(), `{"success": "ping"}`)
		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("test server ping handler", func(t *testing.T) {
		s := NewServer(":2000", nil).route()
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, "/", nil)

		s.ServeHTTP(w, r)

		assert.Equal(t, w.Body.String(), "404 page not found\n")
		assert.Equal(t, w.Code, http.StatusNotFound)
	})
}
