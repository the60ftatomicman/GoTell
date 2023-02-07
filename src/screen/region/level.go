package region

import (
	"example/gotell/src/tile"
)

const MAP_LEFT = 0
const MAP_TOP = 0
const MAP_LINES = 19
const MAP_COLUMNS = 80

type Level struct {
	Name   string
	Filename string
	Buffer [][]tile.Tile
}

func (m *Level) Initialize(b [][]tile.Tile) {
	m.Buffer  = initializeBuffer(MAP_LINES, MAP_COLUMNS, b,tile.FOG)
}

func (m *Level) Get() (int, int, int, int, [][]tile.Tile) {
	return MAP_LEFT, MAP_TOP, MAP_LINES, MAP_COLUMNS, m.Buffer
}

func (m *Level) Refresh() () {
	//TODO -- maybe move some logic into here? Not used in this region tbh
}

func (m *Level) ReadDataFromFile() [][]tile.Tile {
	tiles := [][]tile.Tile{}
	fileData := []string{
		"b", // aLwAYS need to include this. idk why.
		"80w",
		"w",
		"w",
		"w",
		"w",
		"w",
		"w",
		"5w",
		"5w",
		"5w",
		"5w",
		//BUG - adding fog for this is bizarre revisit that code
		"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
		//"10w,20b,20w,20b,10w",
	}
	for _,row := range fileData {
		tiles = append(tiles,tile.FileParser(row))
	}
	return tiles
}