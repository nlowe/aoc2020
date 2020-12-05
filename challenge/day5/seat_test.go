package day5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeatID(t *testing.T) {
	tests := []struct {
		assignment string
		id         int
	}{
		{assignment: "FBFBBFFRLR", id: 357},
		{assignment: "BFFFBBFRRR", id: 567},
		{assignment: "FFFBBBFRRR", id: 119},
		{assignment: "BBFFBBFRLL", id: 820},
	}
	for _, tt := range tests {
		t.Run(tt.assignment, func(t *testing.T) {
			require.Equal(t, tt.id, seatID(tt.assignment))
		})
	}

	t.Run("Panics on invalid token", func(t *testing.T) {
		require.PanicsWithError(t, "unknown token parsing assignment abc: a", func() {
			_ = seatID("abc")
		})
	})
}
