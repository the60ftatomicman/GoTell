package region

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strings"
)

const PROFILE_LEFT = 80
const PROFILE_TOP = 0
const PROFILE_LINES = 19
const PROFILE_COLUMNS = 18
// IMPORTANT LINES IN THE PROFILE FOR WRITING!
// START INDEX 0
const LINE_VAR_NAME = 1
const LINE_VAR_CLASS = 2
const LINE_VAR_HEALTH = 4
const LINE_VAR_MANA = 5
//const LINE_VAR_GOLD = 6
const LINE_VAR_LEVEL = 6
const LINE_LBL_ITEMS = 8
// Remember, each of these only have 16 characters!
const LINE_VAR_ITEM_1 = 9
const LINE_VAR_ITEM_2 = 10

//TODO -- just put a player reference in here....
type Profile struct {
	Name   string
	Class  string
	Health string
	Mana   string
	Gold   string
	Level  string
	XP     string
	Items  []string
	SelectedItem string
	Buffer [][]tile.Tile
}

func (p *Profile) Initialize(b [][]tile.Tile) {

	//haha yeah I am going to ignore B passed in.
	b = p.compile()

	p.Buffer = initializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, b,tile.BLANK)
}

func (p *Profile) Get() (int, int, int, int, [][]tile.Tile) {
	return PROFILE_LEFT, PROFILE_TOP, PROFILE_LINES, PROFILE_COLUMNS, p.Buffer
}

func (p *Profile) ReadDataFromFile() [][]tile.Tile {
	// After reading file we ought to see these things
	p.Name   = "Hero"
	p.Class  = "Warrior"
	p.Health = "100"
	p.Mana   = "100"
	p.Gold   = "0"
	p.Level  = "1"
	p.XP     = "0"
	p.Items  = []string{"Boots","Helmet",""}
	
	return [][]tile.Tile{}
}

func (p *Profile)Refresh(){
	p.Buffer = initializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, p.compile(),tile.BLANK)
}

func (p *Profile)compile()[][]tile.Tile{
	t := [][]tile.Tile{{tile.BLANK}}
	t = append(t, tile.GenerateHorizontalDivider(PROFILE_COLUMNS-2,tile.BLANK,tile.PROFILE_H))
	for i := 0; i < PROFILE_LINES-3; i++ {
		switch(i){
			case LINE_VAR_NAME:{
				t = append(t, p.getBaseRow(1," NAME: "+p.Name,core.FgWhite))
			}
			case LINE_VAR_CLASS:{
				t = append(t, p.getBaseRow(1,"CLASS: "+p.Class,core.FgWhite))
			}
			case LINE_VAR_HEALTH:{
				t = append(t, p.getBaseRow(1,"   HP: "+p.Health,core.FgRed))
			}
			case LINE_VAR_MANA:{
				t = append(t, p.getBaseRow(1," MANA: "+p.Mana,core.FgBlue))
			}
			case LINE_VAR_LEVEL:{
				t = append(t, p.getBaseRow(1,"LEVEL: "+p.Level+" XP: "+p.XP,core.FgYellow))
			}
			//case LINE_VAR_GOLD:{
			//	t = append(t, p.getBaseRow(1,"GOLD: "+p.Gold,core.FgYellow))
			//}
			case LINE_LBL_ITEMS:{
				t = append(t, p.getBaseRow(0," --- ITEMS --- ",core.FgCyan))
			}
			case LINE_VAR_ITEM_1:{
				if(p.SelectedItem == "1"){
					t = append(t, p.getBaseRow(1,"1) "+p.Items[0],core.FgWhite,core.BgBlue))
				}else{
					t = append(t, p.getBaseRow(1,"1) "+p.Items[0],core.FgBlack))
				}
			}
			case LINE_VAR_ITEM_2:{
				if(p.SelectedItem == "2"){
					t = append(t, p.getBaseRow(1,"2) "+p.Items[1],core.FgWhite,core.BgBlue))
				}else{
					t = append(t, p.getBaseRow(1,"2) "+p.Items[1],core.FgBlack))
				}
			}
			default:{t = append(t, p.getBaseRow(0,"",core.FgBlue))}
		}
		
	}
	t = append(t, tile.GenerateHorizontalDivider(PROFILE_COLUMNS-2,tile.BLANK,tile.PROFILE_H))
	return t
}

func (p *Profile)getBaseRow(colIdx int, extraMsg string,colors ...core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{tile.PROFILE_V}
	msgArray := strings.Split(extraMsg, "")
	endIdx   := colIdx+len(msgArray)

	if(endIdx > PROFILE_COLUMNS-2){
		endIdx = PROFILE_COLUMNS-2
	}

	var bgColor core.TermCodes = core.BgGrey
	if(len(colors) > 1){
		bgColor = colors[1]
	}

	for i := 0; i < PROFILE_COLUMNS-2; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],colors[0],bgColor))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",colors[0],core.BgGrey))
		}
	}
	t = append(t, tile.PROFILE_V)
	return t
}

