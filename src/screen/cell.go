package screen

import (
	"example/gotell/src/tile"
)

// Cell
// An array pointer of tiles so we can place
// Tiles on the Z axis
type Cell struct {
	Tiles []tile.Tile
}

// generateNewCell
// Used to create a new cell and populates it with tile.Blank
func generateNewCell() Cell{
	c := Cell{}
	c.Tiles = []tile.Tile{tile.BLANK}
	return c 
}

// Get
// Returns the last tile in a cell
func (c *Cell)Get() tile.Tile {
	return c.Tiles[len(c.Tiles)-1]
}

// Set
// Push a tile to the end of the cell
func (c *Cell)Set(t tile.Tile) {
	c.Tiles = append(c.Tiles,t);
}

// Pop
// Removes the LAST entry in the cell.
// If nothing is present after that we populate the cell with a BLANK
func (c *Cell)Pop() {
	c.Tiles = c.Tiles[:len(c.Tiles)-1]
	if (len(c.Tiles) == 0){
		c.Set(tile.BLANK)
	}
}