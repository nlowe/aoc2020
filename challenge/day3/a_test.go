package day3

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
	"github.com/stretchr/testify/require"
)

const day3example = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(day3example)

	result := partA(input)

	require.Equal(t, 7, result)
}
