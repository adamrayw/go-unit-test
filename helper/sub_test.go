package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		firstname := "Adam"
		assert.Equal(t, "Adam", firstname, "Result must be Adam")
	})

	t.Run("Test2", func(t *testing.T) {
		lastname := "Ray"
		assert.Equal(t, "Rays", lastname, "Result must be Ray")
	})
}
