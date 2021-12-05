package day23

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	t.Skipf("Challenge not yet solved")
	input := challenge.FromLiteral("foobar")

	result := partB(input)

	require.Equal(t, 42, result)
}
