package day21

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := partB(input)

	require.Equal(t, 42, result)
}
