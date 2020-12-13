package day13

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day13example = `939
7,13,x,x,59,x,31,19`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day13example)

	result := partA(input)

	require.Equal(t, 295, result)
}
