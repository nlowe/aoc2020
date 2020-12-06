package day6

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day6example = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day6example)

	result := partA(input)

	require.Equal(t, 11, result)
}
