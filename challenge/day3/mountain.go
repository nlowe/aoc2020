package day3

import "github.com/nlowe/aoc2020/challenge"

const (
	tileOpen = '.'
	tileTree = '#'
)

type mountain struct {
	m []rune

	w int
	h int
}

func parseMountain(input *challenge.Input) mountain {
	lines := input.LineSlice()

	result := mountain{
		m: make([]rune, len(lines)*len(lines[0])),
		w: len(lines[0]),
		h: len(lines),
	}

	for row, line := range lines {
		for column, tile := range line {
			result.m[result.w*row+column] = tile
		}
	}

	return result
}

func (m mountain) tileAt(x, y int) rune {
	return m.m[(m.w*y)+(x%m.w)]
}

func (m mountain) treesAlongSlope(dx, dy int) int {
	trees := 0

	x := 0
	y := 0
	for y < m.h {
		if m.tileAt(x, y) == tileTree {
			trees++
		}

		x += dx
		y += dy
	}

	return trees
}
