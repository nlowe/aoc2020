package day2

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day2example = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day2example)

	result := a(input)

	require.Equal(t, 2, result)
}
