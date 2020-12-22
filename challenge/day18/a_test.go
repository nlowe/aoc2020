package day18

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	tests := []struct {
		eqn      string
		expected int
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}
	for _, tt := range tests {
		t.Run(tt.eqn, func(t *testing.T) {
			input := challenge.FromLiteral(tt.eqn)

			result := partA(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
