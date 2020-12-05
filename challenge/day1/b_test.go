package day1

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(day1example)

	result := partB(input)

	require.Equal(t, 241861950, result)
}
