package day10

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day10exampleA = `16
10
15
5
1
11
7
19
6
12
4`

const day10exampleb = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestA(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{day10exampleA, 35},
		{day10exampleb, 220},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			input := challenge.FromLiteral(tt.input)

			result := partA(input)

			require.Equal(t, tt.expected, result)
		})
	}
}
