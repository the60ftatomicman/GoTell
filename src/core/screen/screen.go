package screen

import (
	"example/gotell/src/core"
	"example/gotell/src/core/tile"
	"fmt"
	"strconv"
)

//TODO 3/8/2023 -- add initializer to this class and allow for dynamic WIDTH and HEIGHT. See popup region for details
//TODO 6/9/2023 -- why did I make this width and height, not column and lines?

const SCREEN_WIDTH = 100 // How many COLUMNS of characters the screen is
const SCREEN_HEIGHT = 29 // How many LINES of characters the screen is

// --------------------

// Screen
// This object helps manage the data which we want to draw to the screen
// The game logic should interact with the buffer and the Raw property
// is the compiled text to send to the client
type Screen struct {
	Buffer [SCREEN_HEIGHT][SCREEN_WIDTH] tile.Cell    // Collection of Cells containing Z amount of tiles
	Raw    string                        `default:""` // Represents the raw ANSI text we wish to send to the terminal
}

// BlankScreen
// Set every cell to just contain tile.Blank
func BlankScreen() [SCREEN_HEIGHT][SCREEN_WIDTH] tile.Cell {
	var blank [SCREEN_HEIGHT][SCREEN_WIDTH]tile.Cell
	for l := 0; l < len(blank); l++ {
		for c := 0; c < len(blank[l]); c++ {
			blank[l][c] = tile.GenerateNewCell()
		}
	}
	return blank
}

// Compile
// When the regions are modified and we need to modify the cells
// This functino takes the tiles from those regions and pushes them into the Cell which is a slice.
// Order of Operations matter! Latter passed in regions display on top of earlier regions.
// TODO -- should I blank before I run this each time?
func (s *Screen) Compile(regionList ...IRegion) {
	s.Raw = ""
	for _, r := range regionList {
		rLeft, rTop, rLines, rColumns, rBuffer := r.Get()
		for l := 0; l < rLines; l++ {
			for c := 0; c < rColumns; c++ {
				//s.Buffer[rTop+l][rLeft+c].Tiles[0] = rBuffer[l][c] //TODO -- this works but sadly we need to reconsider using SET
				s.Buffer[rTop+l][rLeft+c].Set(rBuffer[l][c])
			}
		}
	}
}

// Refresh
// This resets Raw based on what is currently in our Buffer
func (s *Screen) Refresh() {
	s.Raw = ""
	for _, line := range s.Buffer {
		for _, column := range line {
			if column.Get().BGColor != "" {
				s.Raw += core.GenChar(string(column.Get().Icon), column.Get().Color, column.Get().BGColor)
			} else {
				s.Raw += core.GenChar(string(column.Get().Icon), column.Get().Color)
			}
		}
		s.Raw += string(core.TermCodes(core.Newline))
	}
}

// Set
// Attempts to put a tile into a Cell slice
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

// Get
// Returns the Raw text
func (s *Screen) Get() string {
	return s.Raw
}
