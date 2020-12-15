package day15

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected int
	}{
		{"0,3,6", 175594},
		{"1,3,2", 2578},
		{"2,1,3", 3544142},
		{"1,2,3", 261214},
		{"2,3,1", 6895259},
		{"3,2,1", 18},
		{"3,1,2", 362},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			input := challenge.FromLiteral(tt.input)

			result := partB(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
