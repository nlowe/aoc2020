package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestIntMin(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 0, b: -1, expected: -1},
		{a: -1, b: 0, expected: -1},
		{a: 0, b: 1, expected: 0},
		{a: 1, b: 0, expected: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.a, tt.b), func(t *testing.T) {
			assert.Equal(t, tt.expected, IntMin(tt.a, tt.b))
		})
	}
}

func TestIntMax(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 0, b: -1, expected: 0},
		{a: -1, b: 0, expected: 0},
		{a: 0, b: 1, expected: 1},
		{a: 1, b: 0, expected: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.a, tt.b), func(t *testing.T) {
			assert.Equal(t, tt.expected, IntMax(tt.a, tt.b))
		})
	}
}

func TestIntClamp(t *testing.T) {
	tests := []struct {
		low      int
		n        int
		high     int
		expected int
	}{
		{low: 0, n: 1, high: 2, expected: 1},
		{low: 0, n: -1, high: 2, expected: 0},
		{low: 0, n: 3, high: 2, expected: 2},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d <= %d <= %d", tt.low, tt.n, tt.high), func(t *testing.T) {
			assert.Equal(t, tt.expected, IntClamp(tt.low, tt.n, tt.high))
		})
	}

	t.Run("Panics if low > high", func(t *testing.T) {
		assert.PanicsWithError(t, "IntClamp: low cannot be > high: 2 > 1", func() {
			_ = IntClamp(2, 0, 1)
		})
	})
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		x1       int
		y1       int
		x2       int
		y2       int
		expected int
	}{
		{2, 3, 4, 5, 4},
		{2, 3, 4, -5, 10},
		{2, 3, -4, 5, 8},
		{2, 3, -4, -5, 14},
		{2, -3, 4, 5, 10},
		{2, -3, 4, -5, 4},
		{2, -3, -4, 5, 14},
		{2, -3, -4, -5, 8},
		{-2, 3, 4, 5, 8},
		{-2, 3, 4, -5, 14},
		{-2, 3, -4, 5, 4},
		{-2, 3, -4, -5, 10},
		{-2, -3, 4, 5, 14},
		{-2, -3, 4, -5, 8},
		{-2, -3, -4, 5, 10},
		{-2, -3, -4, -5, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d,%d,%d", tt.x1, tt.y1, tt.x2, tt.y2), func(t *testing.T) {
			require.Equal(t, tt.expected, ManhattanDistance(tt.x1, tt.y1, tt.x2, tt.y2))
		})
	}
}
