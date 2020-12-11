package day11

import "github.com/nlowe/aoc2020/challenge"

const (
	tileFloor        = '.'
	tileEmptySeat    = 'L'
	tileOccupiedSeat = '#'
)

type (
	neighborFunc func(*challenge.TileMap, int, int) int
	raycastFunc  func(seats *challenge.TileMap, x, y, dx, dy int) (rune, bool)
)

// step progresses the simulation by one iteration, finding neighbors using the specified
// function and detecting suitable seats with the specified threshold. It returns true
// iff any seats changed this iteration.
func step(seats *challenge.TileMap, neighbors neighborFunc, threshold int) bool {
	w, h := seats.Size()
	work := challenge.NewTileMap(w, h)

	changed := false
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			neighborCount := neighbors(seats, x, y)
			tile, _ := seats.TileAt(x, y)

			switch {
			case tile == tileEmptySeat && neighborCount == 0:
				changed = true
				work.SetTile(x, y, tileOccupiedSeat)
			case tile == tileOccupiedSeat && neighborCount >= threshold:
				changed = true
				work.SetTile(x, y, tileEmptySeat)
			default:
				work.SetTile(x, y, tile)
			}
		}
	}

	*seats = *work
	return changed
}

func countOccupiedNeighbors(seats *challenge.TileMap, x, y int, search raycastFunc) int {
	count := 0

	d := []struct {
		x int
		y int
	}{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}
	for _, delta := range d {
		if tile, ok := search(seats, x, y, delta.x, delta.y); ok && tile == tileOccupiedSeat {
			count++
		}
	}

	return count
}

func countOccupied(seats *challenge.TileMap) int {
	count := 0

	w, h := seats.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			if tile, _ := seats.TileAt(x, y); tile == tileOccupiedSeat {
				count++
			}
		}
	}

	return count
}
