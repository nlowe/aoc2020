package day8

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day8example = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day8example)

	result := partA(input)

	require.Equal(t, 5, result)
}
