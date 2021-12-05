package day25

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	t.Skipf("Challenge not yet solved")
	input := challenge.FromLiteral("foobar")

	result := partA(input)

	require.Equal(t, 42, result)
}
