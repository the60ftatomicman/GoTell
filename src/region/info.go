package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strings"
)

const INFO_LEFT = 1
const INFO_TOP = 23
const INFO_LINES = 6
const INFO_COLUMNS = 79

// Info
// Displays on the bottom beneath the level.
// Provides current commands available (non menu) and other info
type Info struct {
	Message [INFO_LINES-2]string
	Buffer  [][]tile.Tile
}

func (p *Info) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	p.Set("By Andrew Garber")
	b = p.compile()

	p.Buffer = screen.InitializeBuffer(INFO_LINES, INFO_COLUMNS, b,tile.BLANK)
}

func (p *Info) Refresh() {
	b := p.compile()
	p.Buffer = screen.InitializeBuffer(INFO_LINES, INFO_COLUMNS, b,tile.BLANK) // rename to generateBuffer?
}

func (p *Info) Get() (int, int, int, int, [][]tile.Tile) {
	return INFO_LEFT, INFO_TOP, INFO_LINES, INFO_COLUMNS, p.Buffer
}

func (p *Info) Set(msgs ...string){
	for i:=0; i<INFO_LINES-2;i++ {
		if (i < len(msgs)){
			p.Message[i] = msgs[i]
		}else{
			p.Message[i] = " "
		}
	}
}

func (p *Info) compile()[][]tile.Tile{
	t := [][]tile.Tile{tile.GenerateHorizontalDivider(INFO_COLUMNS-2,tile.BLANK,overrides.INFO_H)}
	for i := 0; i < len(p.Message); i++ {
		t = append(t, p.getBaseRow(1,p.Message[i],core.FgWhite))
	}
	t = append(t, tile.GenerateHorizontalDivider(INFO_COLUMNS-2,tile.BLANK,overrides.INFO_H))
	return t
}

func (p *Info) getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{overrides.INFO_V}
	msgArray := strings.Split(extraMsg, "")
	endIdx   := colIdx+len(msgArray)

	if(endIdx > INFO_COLUMNS-2){
		endIdx = INFO_COLUMNS-2
	}

	for i := 0; i < INFO_COLUMNS-2; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],color,core.BgBlack))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",color,core.BgBlack))
		}
	}
	t = append(t, overrides.INFO_V)
	return t
}
