package day13

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{day13example, "1068781"},
		{`0
17,x,13,19`, "3417"},
		{`0
67,7,59,61`, "754018"},
		{`0
67,x,7,59,61`, "779210"},
		{`0
67,7,x,59,61`, "1261476"},
		{`0
1789,37,47,1889`, "1202161486"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := partB(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
