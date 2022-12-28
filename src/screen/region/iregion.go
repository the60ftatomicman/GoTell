package region

import "example/gotell/src/tile"

type IRegion interface {
	Get() (int, int, int, int, [][]tile.Tile)
	Initialize(b [][]tile.Tile)
}

func initializeBuffer(lines int, columns int, data [][]tile.Tile) [][]tile.Tile {
	buffer := [][]tile.Tile{}
	for l := 0; l < lines; l++ {
		tileLine := []tile.Tile{}
		for c := 0; c < columns; c++ {
			if l < len(data) && c < len(data[l]) {
				tileLine = append(tileLine, data[l][c])
			} else {
				tileLine = append(tileLine, tile.BLANK)
			}
		}
		buffer = append(buffer, tileLine)
	}
	return buffer
}
