package server

import (
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
