package screen

import (
	"example/gotell/src/tile"
)

// An array pointer of tiles so we can place
// Tiles on the Z axis
type Cell struct {
	Tiles []tile.Tile
}

func generateNewCell() Cell{
	c := Cell{}
	c.Tiles = []tile.Tile{tile.BLANK}
	return c 
}
func (c *Cell)Get() tile.Tile {
	return c.Tiles[len(c.Tiles)-1]
}

func (c *Cell)Set(t tile.Tile) {
	c.Tiles = append(c.Tiles,t);
}

func (c *Cell)Pop() {
	c.Tiles = c.Tiles[:len(c.Tiles)-1]
}