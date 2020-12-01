package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustAtoI(t *testing.T) {
	t.Run("Parses", func(t *testing.T) {
		assert.Equal(t, 123, MustAtoI("123"))
	})

	t.Run("Panics on Error", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = MustAtoI("abc")
		})
	})
}
