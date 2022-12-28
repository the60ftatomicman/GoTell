package region

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strings"
)

const INFO_LEFT = 1
const INFO_TOP = 21
const INFO_LINES = 4
const INFO_COLUMNS = 99
// IMPORTANT LINES IN THE INFO SECTION FOR WRITING!
// START INDEX 0
const LINE_VAR_INFO_MSG = 0

type Info struct {
	Message string
	Buffer  [][]tile.Tile
}

func (p *Info) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	b = p.populate()

	p.Buffer = initializeBuffer(INFO_LINES, INFO_COLUMNS, b)
}

func (p *Info) Get() (int, int, int, int, [][]tile.Tile) {
	return INFO_LEFT, INFO_TOP, INFO_LINES, INFO_COLUMNS, p.Buffer
}

func (p *Info)populate()[][]tile.Tile{
	t := [][]tile.Tile{{tile.BLANK}}
	t = append(t, getHorizontalDivider_INFO())
	for i := 0; i < INFO_LINES-3; i++ {
		switch(i){
			case LINE_VAR_INFO_MSG:{
				t = append(t, getBaseRow_INFO(1,p.Message,core.FgWhite))
			}
			default:{t = append(t, getBaseRow_INFO(1," ",core.FgBlack))}
		}
		
	}
	t = append(t, getHorizontalDivider_INFO())
	return t
}

//TODO make part of struct
func getBaseRow_INFO(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{tile.INFO_V}
	msgArray := strings.Split(extraMsg, "")
	endIdx   := colIdx+len(msgArray)

	if(endIdx > INFO_COLUMNS-2){
		endIdx = INFO_COLUMNS-2
	}

	for i := 0; i < INFO_COLUMNS-2; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],color,core.BgGrey))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",color,core.BgGrey))
		}
	}
	t = append(t, tile.INFO_V)
	return t
}
//TODO make part of struct
func getHorizontalDivider_INFO() []tile.Tile {
	t := []tile.Tile{tile.BLANK}
	for i := 0; i < INFO_COLUMNS-2; i++ {
		t = append(t, tile.INFO_H)
	}
	t = append(t, tile.BLANK)
	return t
}

