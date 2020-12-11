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
		{`10
6
4
7
1
5`, 4},
		{`4
11
7
8
1
6
5`, 7},
		{`3
1
6
2`, 4},
		{`17
6
10
5
13
7
1
4
12
11
14`, 28},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := partB(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
