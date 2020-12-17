package day17

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day17example = `.#.
..#
###`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day17example)

	result := partA(input)

	require.Equal(t, 112, result)
}
