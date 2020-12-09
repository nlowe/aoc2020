package day9

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(day9example)

	// This is the best proxy we get, the example input from today
	// uses a different preamble length than our input
	result := findWeakness(input, 5)

	require.Equal(t, 62, result)
}
