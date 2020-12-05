package day5

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`)

	result := partA(input)

	require.Equal(t, 820, result)
}
