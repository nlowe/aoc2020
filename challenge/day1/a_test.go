package day1

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day1example = `1721
979
366
299
675
1456
`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day1example)

	result := a(input)

	require.Equal(t, 514579, result)
}
