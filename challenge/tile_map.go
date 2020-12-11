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

func (t *TileMap) outOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= t.w || y >= t.h
}

func (t *TileMap) indexOf(x, y int) (int, bool) {
	return x + (t.w * y), !t.outOfBounds(x, y)
}

func (t *TileMap) SetTile(x, y int, tile rune) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		panic(fmt.Errorf("out of bounds tile access: [%d, %d] is not within the %dx%d map", x, y, t.w, t.h))
	}

	t.tiles[idx] = tile
}

func (t *TileMap) TileAt(x, y int) (rune, bool) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		return ' ', false
	}

	return t.tiles[idx], true
}
