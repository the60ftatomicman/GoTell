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
	Buffer [SCREEN_HEIGHT][SCREEN_WIDTH]tile.Tile
	Raw    string
}

func BlankScreen() [SCREEN_HEIGHT][SCREEN_WIDTH]tile.Tile {
	var blank [SCREEN_HEIGHT][SCREEN_WIDTH]tile.Tile
	for lIdx, line := range blank {
		for cIdx, _ := range line {
			blank[lIdx][cIdx] = tile.BLANK
		}
	}
	return blank
}
func (s *Screen) Get() string {
	return s.Raw
}

func (s *Screen) Compile(regionList ...region.IRegion) {
	s.Raw = ""
	for _, r := range regionList {
		rLeft, rTop, rLines, rColumns, rBuffer := r.Get()
		for l := 0; l < rLines; l++ {
			for c := 0; c < rColumns; c++ {
				if s.Buffer[rTop+l][rLeft+c].Name != "PLAYER" && 
				   s.Buffer[rTop+l][rLeft+c].Name != "ENEMY" {
					s.Buffer[rTop+l][rLeft+c] = rBuffer[l][c]
				}
			}
		}
	}
	//add buffer
	for _, line := range s.Buffer {
		for _, column := range line {
			if (column.BGColor != ""){
				s.Raw += core.GenChar(string(column.Icon), column.Color,column.BGColor) //TODO -- rename!
			}else{
				s.Raw += core.GenChar(string(column.Icon), column.Color)
			}
		}
		s.Raw += string(core.TermCodes(core.Newline))
	}
}
func (s *Screen) Set(t tile.Tile, idx ...int) {
	var line = idx[1]
	var column = idx[0]

	if line >= SCREEN_HEIGHT {
		fmt.Println("ERROR, line " + strconv.Itoa(line) + " is greater than " + strconv.Itoa(SCREEN_HEIGHT))
	} else {
		if column >= SCREEN_WIDTH {
			fmt.Println("ERROR, column " + strconv.Itoa(column) + " is greater than " + strconv.Itoa(SCREEN_WIDTH))
		} else {
			s.Buffer[line][column] = t
		}
	}
}
