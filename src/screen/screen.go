package screen

import (
	"example/gotell/src/core"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"fmt"
	"strconv"
)

const SCREEN_WIDTH = 100
const SCREEN_HEIGHT = 25

type Screen struct {
	Buffer [SCREEN_HEIGHT][SCREEN_WIDTH]Cell
	Raw    string
}

func BlankScreen() [SCREEN_HEIGHT][SCREEN_WIDTH]Cell {
	var blank [SCREEN_HEIGHT][SCREEN_WIDTH]Cell
	for l := 0; l < len(blank); l++ {
		for c := 0; c < len(blank[l]); c++ {
			blank[l][c] = generateNewCell()
		}
	}
	return blank
}

//When the regions are modified and we need to make new cells; call this
func (s *Screen) Compile(regionList ...region.IRegion) {
	s.Raw = ""
	for _, r := range regionList {
		rLeft, rTop, rLines, rColumns, rBuffer := r.Get()
		for l := 0; l < rLines; l++ {
			for c := 0; c < rColumns; c++ {
				s.Buffer[rTop+l][rLeft+c].Tiles[0] = rBuffer[l][c]
			}
		}
	}
}

//When we are just updating what our cells look like; call this.
func (s *Screen) Refresh() {
	s.Raw = ""
	for _, line := range s.Buffer {
		for _, column := range line {
			if (column.Get().BGColor != ""){
				s.Raw += core.GenChar(string(column.Get().Icon), column.Get().Color,column.Get().BGColor)
			}else{
				s.Raw += core.GenChar(string(column.Get().Icon), column.Get().Color)
			}
		}
		s.Raw += string(core.TermCodes(core.Newline))
	}
}

func (s *Screen) Set(t tile.Tile, idx ...int) {
	var column = idx[1]
	var line = idx[0]

	if line >= SCREEN_HEIGHT {
		fmt.Println("ERROR, line " + strconv.Itoa(line) + " is greater than " + strconv.Itoa(SCREEN_HEIGHT))
	} else {
		if column >= SCREEN_WIDTH {
			fmt.Println("ERROR, column " + strconv.Itoa(column) + " is greater than " + strconv.Itoa(SCREEN_WIDTH))
		} else {
			s.Buffer[line][column].Set(t)
		}
	}
}

func (s *Screen) Get() string {
	return s.Raw
}
