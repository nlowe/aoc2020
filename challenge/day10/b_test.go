package day10

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{day10exampleA, 8},
		{day10exampleb, 19208},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := partB(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
