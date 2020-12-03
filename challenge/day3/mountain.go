package day3

import "github.com/nlowe/aoc2020/challenge"

const tileTree = '#'

// mountain is a simple wrapper around challenge.TileMap that
// scales infinitely in the X direction.
type mountain struct {
	*challenge.TileMap
}

// TileAt wraps (challenge.TileMap).TileAt(...) as if the map
// repeats infinitely in the X direction.
func (m mountain) TileAt(x, y int) rune {
	w, _ := m.Size()

	return m.TileMap.TileAt(x%w, y)
}

func (m mountain) treesAlongSlope(dx, dy int) int {
	_, h := m.Size()

	trees := 0

	x := 0
	y := 0
	for y < h {
		if m.TileAt(x, y) == tileTree {
			trees++
		}

		x += dx
		y += dy
	}

	return trees
}
