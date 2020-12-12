package day12

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day12example = `F10
N3
F7
R90
F11`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day12example)

	result := partA(input)

	require.Equal(t, 25, result)
}
