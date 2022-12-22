package region

import (
	"example/gotell/src/tile"
)

const PROFILE_LEFT = 80
const PROFILE_TOP = 0
const PROFILE_LINES = 20
const PROFILE_COLUMNS = 18

type Profile struct {
	Name   string
	Buffer [][]tile.Tile
}

func (p *Profile) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	b = [][]tile.Tile{
		{tile.BLANK},
		getHorizontalDiverder(),
		{tile.PROFILE_V},
		{tile.PROFILE_V},
		{tile.PROFILE_V},
		{tile.PROFILE_V},
		{tile.PROFILE_V},
		{tile.PROFILE_V},
	}

	p.Buffer = initializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, b)
}

func (p *Profile) Get() (int, int, int, int, [][]tile.Tile) {
	return PROFILE_LEFT, PROFILE_TOP, PROFILE_LINES, PROFILE_COLUMNS, p.Buffer
}

func getHorizontalDiverder() []tile.Tile {
	t := []tile.Tile{}
	for i := 0; i < PROFILE_COLUMNS; i++ {
		t = append(t, tile.PROFILE_V)
	}
	return t
}
