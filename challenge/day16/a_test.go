package day16

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day16example = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day16example)

	result := partA(input)

	require.Equal(t, 71, result)
}
