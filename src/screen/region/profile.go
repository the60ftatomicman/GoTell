package region

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strings"
)

const PROFILE_LEFT = 80
const PROFILE_TOP = 0
const PROFILE_LINES = 20
const PROFILE_COLUMNS = 18
// IMPORTANT LINES IN THE PROFILE FOR WRITING!
// START INDEX 0
const LINE_VAR_NAME = 1
const LINE_VAR_CLASS = 2
const LINE_VAR_HEALTH = 4
const LINE_VAR_MANA = 5
const LINE_VAR_GOLD = 6

const LINE_LBL_ITEMS = 8
// Remember, each of these only have 16 characters!
const LINE_VAR_ITEM_1 = 9
const LINE_VAR_ITEM_2 = 10

type Profile struct {
	Name   string
	Class  string
	Health string
	Mana   string
	Gold   string
	Items  []string
	Buffer [][]tile.Tile
}

func (p *Profile) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	b = p.populate()

	p.Buffer = initializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, b)
}

func (p *Profile) Get() (int, int, int, int, [][]tile.Tile) {
	return PROFILE_LEFT, PROFILE_TOP, PROFILE_LINES, PROFILE_COLUMNS, p.Buffer
}

func (p *Profile)populate()[][]tile.Tile{
	t := [][]tile.Tile{{tile.BLANK}}
	t = append(t, getHorizontalDivider())
	for i := 0; i < PROFILE_LINES-3; i++ {
		switch(i){
			case LINE_VAR_NAME:{
				t = append(t, getBaseRow(1," NAME: "+p.Name,core.FgWhite))
			}
			case LINE_VAR_CLASS:{
				t = append(t, getBaseRow(1,"CLASS: "+p.Class,core.FgWhite))
			}
			case LINE_VAR_HEALTH:{
				t = append(t, getBaseRow(1,"  HP: "+p.Health,core.FgRed))
			}
			case LINE_VAR_MANA:{
				t = append(t, getBaseRow(1,"MANA: "+p.Mana,core.FgBlue))
			}
			case LINE_VAR_GOLD:{
				t = append(t, getBaseRow(1,"GOLD: "+p.Gold,core.FgYellow))
			}
			case LINE_LBL_ITEMS:{
				t = append(t, getBaseRow(0," --- ITEMS --- ",core.FgCyan))
			}
			case LINE_VAR_ITEM_1:{
				t = append(t, getBaseRow(1,"1) "+p.Items[0],core.FgMagenta))
			}
			case LINE_VAR_ITEM_2:{
				t = append(t, getBaseRow(1,"2) "+p.Items[1],core.FgMagenta))
			}
			default:{t = append(t, getBaseRow(0,"",core.FgBlue))}
		}
		
	}
	t = append(t, getHorizontalDivider())
	return t
}

func getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{tile.PROFILE_V}
	msgArray := strings.Split(extraMsg, "")
	endIdx   := colIdx+len(msgArray)

	if(endIdx > PROFILE_COLUMNS-2){
		endIdx = PROFILE_COLUMNS-2
	}

	for i := 0; i < PROFILE_COLUMNS-2; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.PROFILE_TEXT(msgArray[i-colIdx],color))
		}else{
			t = append(t, tile.BLANK)
		}
	}
	t = append(t, tile.PROFILE_V)
	return t
}

func getHorizontalDivider() []tile.Tile {
	t := []tile.Tile{tile.BLANK}
	for i := 0; i < PROFILE_COLUMNS-2; i++ {
		t = append(t, tile.PROFILE_H)
	}
	t = append(t, tile.BLANK)
	return t
}

