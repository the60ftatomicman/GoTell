package region

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strings"
)

const HEADER_LEFT = 0
const HEADER_TOP = 1
const HEADER_LINES = 2
const HEADER_COLUMNS = 79

type Header struct {
	Message string
	Buffer  [][]tile.Tile
}

func (p *Header) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	p.Set("Some title v0.0.7 - MENUS: (Q)uit (i)ventory (m)ove")
	b = p.compile()

	p.Buffer = initializeBuffer(HEADER_LINES, HEADER_COLUMNS, b,tile.BLANK)
}

func (p *Header) Refresh() {
	b := p.compile()
	p.Buffer = initializeBuffer(HEADER_LINES, HEADER_COLUMNS, b,tile.BLANK) // rename to generateBuffer?
}

func (p *Header) Get() (int, int, int, int, [][]tile.Tile) {
	return HEADER_LEFT, HEADER_TOP, HEADER_LINES, HEADER_COLUMNS, p.Buffer
}
func (p *Header) Set(msgs string){
	p.Message = msgs
}
func (p *Header) compile()[][]tile.Tile{
	t := [][]tile.Tile{p.getBaseRow(1,p.Message,core.FgWhite)}
	t = append(t, tile.GenerateHorizontalDivider(HEADER_COLUMNS,tile.BLANK,tile.INFO_H))
	return t
}

func (p *Header) getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{} //TODO im lazy and just not making tile
	msgArray := strings.Split(extraMsg, "")
	endIdx   := len(msgArray)+1

	if(endIdx > HEADER_COLUMNS){
		endIdx = HEADER_COLUMNS
	}

	for i := 0; i < HEADER_COLUMNS; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],color,core.BgBlack))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",color,core.BgBlack))
		}
	}
	return t
}
