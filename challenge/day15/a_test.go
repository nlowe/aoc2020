package day15

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected int
	}{
		{"0,3,6", 436},
		{"1,3,2", 1},
		{"2,1,3", 10},
		{"1,2,3", 27},
		{"2,3,1", 78},
		{"3,2,1", 438},
		{"3,1,2", 1836},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			input := challenge.FromLiteral(tt.input)

			result := partA(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
