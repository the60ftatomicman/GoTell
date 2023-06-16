package tile

// Cell
// An array pointer of tiles so we can place
// Tiles on the Z axis
type Cell struct {
	Tiles []Tile
}

// generateNewCell
// Used to create a new cell and populates it with a Blank tile
func GenerateNewCell() Cell{
	c := Cell{}
	c.Tiles = []Tile{BLANK}
	return c 
}

// Get
// Returns the last tile in a cell
func (c *Cell)Get() Tile {
	return c.Tiles[len(c.Tiles)-1]
}

// Set
// Push a tile to the end of the cell
func (c *Cell)Set(t Tile) {
	c.Tiles = append(c.Tiles,t);
}

// Pop
// Removes the LAST entry in the cell.
// If nothing is present after that we populate the cell with a BLANK
func (c *Cell)Pop() {
	c.Tiles = c.Tiles[:len(c.Tiles)-1]
	if (len(c.Tiles) == 0){
		c.Set(BLANK)
	}
}

// Clear
// For when you REALLY want to remove everything.
func (c *Cell) Clear() {
	c.Tiles = []Tile{BLANK}
}