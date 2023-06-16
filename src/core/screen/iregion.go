package screen

import "example/gotell/src/core/tile"

// IRegion
// This interface is intended to help build and manage regions of tiles
// to be displayed on the screen
type IRegion interface {
	Get() (int, int, int, int, [][]tile.Tile)
	Initialize(b [][]tile.Tile)
	Refresh()
}

func InitializeBuffer(lines int, columns int, data [][]tile.Tile, defaultTile tile.Tile) [][]tile.Tile{
	buffer := [][]tile.Tile{}
	for l := 0; l < lines; l++ {
		tileLine := []tile.Tile{}
		for c := 0; c < columns; c++ {
			if l < len(data) && c < len(data[l]) {
				tileLine = append(tileLine, data[l][c])
			} else {
				tileLine = append(tileLine, defaultTile)
			}
		}
		buffer = append(buffer, tileLine)
	}
	return buffer
}
