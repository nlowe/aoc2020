package challenge

import "fmt"

// TileMap represents a fixed size grid of runes. The top-left tile is [0,0]
type TileMap struct {
	tiles []rune
	w     int
	h     int
}

func NewTileMap(w, h int) *TileMap {
	return &TileMap{
		tiles: make([]rune, w*h),
		w:     w,
		h:     h,
	}
}

func (t *TileMap) Size() (int, int) {
	return t.w, t.h
}

func (t *TileMap) boundsCheck(x, y int) {
	if x < 0 || y < 0 || x >= t.w || y >= t.h {
		panic(fmt.Errorf("out of bounds tile access: [%d, %d] is not within the %dx%d map", x, y, t.w, t.h))
	}
}

func (t *TileMap) indexOf(x, y int) int {
	t.boundsCheck(x, y)
	return x + (t.w * y)
}

func (t *TileMap) SetTile(x, y int, tile rune) {
	t.tiles[t.indexOf(x, y)] = tile
}

func (t *TileMap) TileAt(x, y int) rune {
	return t.tiles[t.indexOf(x, y)]
}
