package day9

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day9example = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day9example)

	// This is the best proxy we get, the example input from today
	// uses a different preamble length than our input
	_, _, result := findInvalid(input, 5)

	require.Equal(t, 127, result)
}
