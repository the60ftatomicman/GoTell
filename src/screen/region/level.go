package region

import (
	"example/gotell/src/tile"
)

const MAP_LEFT = 0
const MAP_TOP = 0
const MAP_LINES = 20
const MAP_COLUMNS = 80

type Level struct {
	Name   string
	Filename string
	Buffer [][]tile.Tile
	Objects [][]tile.Tile
}

func (m *Level) Initialize(b [][]tile.Tile) {
	m.Buffer  = initializeBuffer(MAP_LINES, MAP_COLUMNS, b)
	m.Objects = initializeBuffer(MAP_LINES, MAP_COLUMNS, b)
}

func (m *Level) Get() (int, int, int, int, [][]tile.Tile) {
	return MAP_LEFT, MAP_TOP, MAP_LINES, MAP_COLUMNS, m.Buffer
}

func (m *Level) ReadDataFromFile() [][]tile.Tile {
	return [][]tile.Tile{
		{tile.BLANK},
		{tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL},
	}
}