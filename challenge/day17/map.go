package day17

import "github.com/nlowe/aoc2020/challenge"

const (
	stateActive   = '#'
	stateInactive = '.'
)

type coord struct {
	x int
	y int
	z int
	w int
}

type dimension struct {
	enableW bool
	tiles   map[coord]rune

	minX int
	maxX int

	minY int
	maxY int

	minZ int
	maxZ int

	minW int
	maxW int
}

func parseDimension(challenge *challenge.Input) *dimension {
	result := &dimension{tiles: map[coord]rune{}}

	for y, row := range challenge.LineSlice() {
		for x, t := range row {
			result.set(x, y, 0, 0, t)
			result.maxX = x
		}
		result.maxY = y
	}

	return result
}

func (d *dimension) set(x, y, z, w int, t rune) {
	d.tiles[coord{x, y, z, w}] = t
}

func (d *dimension) get(x, y, z, w int) rune {
	if t, ok := d.tiles[coord{x, y, z, w}]; ok {
		return t
	}

	return stateInactive
}

func (d *dimension) countNeighbors(x, y, z, w int) int {
	count := 0
	check := func(dx, dy, dz, dw int) {
		if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
			return
		}

		if d.get(x+dx, y+dy, z+dz, w+dw) == stateActive {
			count++
		}
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if d.enableW {
					for dw := -1; dw <= 1; dw++ {
						check(dx, dy, dz, dw)
					}
				} else {
					check(dx, dy, dz, 0)
				}
			}
		}
	}

	return count
}

func (d *dimension) step() {
	work := &dimension{tiles: map[coord]rune{}}

	stepTile := func(x, y, z, w int) {
		t := d.get(x, y, z, w)
		neighborCount := d.countNeighbors(x, y, z, w)

		switch {
		case t == stateActive && (neighborCount == 2 || neighborCount == 3):
			work.set(x, y, z, w, stateActive)
		case t == stateInactive && neighborCount == 3:
			work.set(x, y, z, w, stateActive)
		default:
			work.set(x, y, z, w, stateInactive)
		}
	}

	for x := d.minX - 1; x <= d.maxX+1; x++ {
		for y := d.minY - 1; y <= d.maxY+1; y++ {
			for z := d.minZ - 1; z <= d.maxZ+1; z++ {
				if d.enableW {
					for w := d.minW - 1; w <= d.maxW+1; w++ {
						stepTile(x, y, z, w)
					}
				} else {
					stepTile(x, y, z, 0)
				}
			}
		}
	}

	d.tiles = work.tiles
	d.minX--
	d.minY--
	d.minZ--
	d.minW--
	d.maxX++
	d.maxY++
	d.maxZ++
	d.maxW++
}

func (d *dimension) active() int {
	count := 0
	for _, t := range d.tiles {
		if t == stateActive {
			count++
		}
	}

	return count
}
