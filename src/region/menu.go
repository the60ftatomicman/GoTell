package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	"example/gotell/src/object"
	"strings"
)

const MENU_LEFT    = 0
const MENU_TOP     = 0
const MENU_LINES   = 29
const MENU_COLUMNS = 100
const MENU_COL_DIVIDER = 34
// IMPORTANT LINES IN THE MENU FOR WRITING!
// START INDEX 0
// Remember, each of these only have 16 characters!

const MENU_LINE_VAR_CLASS = 1 // Line for selecting a class
const MENU_LINE_VAR_LEVEL = 3 // Line for selecting the level
var MENU_OPTIONS_CLASS  = []object.PlayerClass{
	object.CLASS_PHYSICAL,
	object.CLASS_MAGIC,
	object.CLASS_SPEED,
	object.CLASS_EXP,
	object.CLASS_FOG,
}
var MENU_OPTIONS_LEVEL  = []string{
	"demolevel",
	"Not Implemented",
} //TODO -- obviously this is getting turned into a struct! 
const MENU_CURSOR_MENU  = 0
const MENU_CURSOR_CLASS = 1
const MENU_CURSOR_LEVEL = 2


// Menu
// Where we are going to setup our character
// This area gives details on what the player is about to suggest.
type Menu struct {
	Player *object.Player
	Cursors []int
	Buffer [][]tile.Tile
}

func (p *Menu) Initialize(b [][]tile.Tile) {
	// 0 == current menu selection
	// 1 == CLASS menu seletion
	// 2 == LEVEL menu selection
	p.Cursors = []int{0,0,0}

	//haha yeah I am going to ignore B passed in.
	b = p.compile()
	p.Buffer = screen.InitializeBuffer(MENU_LINES, MENU_COLUMNS, b,tile.BLANK)
}

func (p *Menu) Get() (int, int, int, int, [][]tile.Tile) {
	return MENU_LEFT, MENU_TOP, MENU_LINES, MENU_COLUMNS, p.Buffer
}

func (p *Menu) updateCursor(cursoridx int, maxLimit int,delta int) {
	p.Cursors[cursoridx] += delta
	if (p.Cursors[cursoridx] < 0){
		p.Cursors[cursoridx] = maxLimit
	} else if (p.Cursors[cursoridx] > maxLimit){
		p.Cursors[cursoridx] = 0
	}
}

func (p *Menu) ChangeSelection(delta int) {
	p.updateCursor(MENU_CURSOR_MENU,1,delta)
}

func (p *Menu) GetSelection() int {
	return p.Cursors[MENU_CURSOR_MENU]
}

func (p *Menu) ChangeClass(delta int) {
	p.updateCursor(MENU_CURSOR_CLASS,len(MENU_OPTIONS_CLASS)-1,delta)
}

func (p *Menu) ChangeLevel(delta int) {
	p.updateCursor(MENU_CURSOR_LEVEL,len(MENU_OPTIONS_LEVEL)-1,delta)
}

func (p *Menu)Refresh(){
	p.Buffer = screen.InitializeBuffer(MENU_LINES, MENU_COLUMNS, p.compile(),tile.BLANK)
}

func (p *Menu)compile()[][]tile.Tile{
	t := [][]tile.Tile{{tile.BLANK}}

	t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
		getBorderTile(),
		getBorderTile(),
	))

	// Draw LEFT side
	t = append(t, p.getBaseRow(1,"----- Choose your settings ----- ",core.FgWhite))
	for i := 0; i < MENU_LINES-4; i++ {
		switch(i){
			case MENU_LINE_VAR_CLASS:{
				fgColor := core.FgWhite
				bgColor := core.BgBlack
				if(p.Cursors[MENU_CURSOR_MENU] == 0){
					fgColor = core.FgBlue
					bgColor = core.BgWhite		
				}
				t = append(t, p.getBaseRow(1," CLASS: "+MENU_OPTIONS_CLASS[p.Cursors[MENU_CURSOR_CLASS]].Name,core.TermCodes(fgColor),core.TermCodes(bgColor)))
				p.Player.Class = MENU_OPTIONS_CLASS[p.Cursors[MENU_CURSOR_CLASS]].Name
				p.Player.Stats = MENU_OPTIONS_CLASS[p.Cursors[MENU_CURSOR_CLASS]].Stats
			}
			case MENU_LINE_VAR_LEVEL:{
				fgColor := core.FgWhite
				bgColor := core.BgBlack
				if(p.Cursors[MENU_CURSOR_MENU] == 1){
					fgColor = core.FgBlue
					bgColor = core.BgWhite		
				}
				t = append(t, p.getBaseRow(1," LEVEL: "+MENU_OPTIONS_LEVEL[p.Cursors[MENU_CURSOR_LEVEL]],core.TermCodes(fgColor),core.TermCodes(bgColor)))
			}
			default: {
				t = append(t, p.getBaseRow(1," ",core.FgWhite,core.BgBlack))
			}
		}
	}
	t[2][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
	t[2] = p.putMessageInTiles(t[2],MENU_COL_DIVIDER+1,"----- Details-----",core.FgWhite)
	t[3][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
	t[3] = p.putMessageInTiles(t[3],MENU_COL_DIVIDER+1,"----- Class",core.FgWhite)
	nextLine:= 4
	// Draw RIGHT side
	classDesc := MENU_OPTIONS_CLASS[p.Cursors[MENU_CURSOR_CLASS]].Description
	for i := 0; i < len(classDesc); i++ {
		idx:= nextLine+i
		t[idx][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
		t[idx] = p.putMessageInTiles(t[idx],MENU_COL_DIVIDER+1,classDesc[i],core.FgGrey)
	}
	nextLine+=len(classDesc)
	t[nextLine][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
	nextLine++
	t[nextLine][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
	t[nextLine] = p.putMessageInTiles(t[nextLine],MENU_COL_DIVIDER+1,"----- Level",core.FgWhite) 
	nextLine++
	levelDesc := []string{"Blah Blah Blah","Some Default Text Here"}//MENU_OPTIONS_CLASS[p.Cursors[MENU_CURSOR_LEVEL]].
	for i := 0; i < len(levelDesc); i++ {
		idx:= nextLine+i
		t[idx][MENU_COL_DIVIDER] = tile.GENERIC_TEXT("|",core.FgCyan,core.BgBlack)
		t[idx] = p.putMessageInTiles(t[idx],MENU_COL_DIVIDER+1,levelDesc[i],core.FgGrey)
	}

	t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
		getBorderTile(),
		getBorderTile(),
	))

	msg := strings.Split("(w/s) change option. (a/d) change selection. Press (g) to begin!","")
	for i:=0; i < len(msg) ; i++ {
		t[len(t)-1][i+1] = tile.GENERIC_TEXT(msg[i],core.TermCodes(core.FgBlack),core.TermCodes(core.BgGreen))
	}
	

	return t
}

func (p *Menu)getBaseRow(colIdx int,extraMsg string,colors ...core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{tile.BLANK}
	msgArray := strings.Split(extraMsg, "")
	endIdx   := colIdx+len(msgArray)

	if(endIdx > MENU_COLUMNS-1){
		endIdx = MENU_COLUMNS-1
	}

	var bgColor core.TermCodes = core.BgBlack
	if(len(colors) > 1){
		bgColor = colors[1]
	}
	//Draw and populate our left Side
	for i := 0; i < MENU_COLUMNS-1; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],colors[0],bgColor))
		}else {
			t = append(t, tile.GENERIC_TEXT(" ",colors[0],core.BgBlack))
		}
	}
	return t
}

func (p *Menu)putMessageInTiles(t []tile.Tile,colIdx int,extraMsg string,colors ...core.TermCodes)[]tile.Tile{
	msgArray := strings.Split(extraMsg, "")
	endIdx   := len(msgArray)

	if(colIdx+endIdx > MENU_COLUMNS-1){
		endIdx = MENU_COLUMNS-1-colIdx
	}

	var bgColor core.TermCodes = core.BgBlack
	if(len(colors) > 1){
		bgColor = colors[1]
	}
	//Draw and populate our left Side
	for i := 0; i < endIdx; i++ {
		t[colIdx+i] = tile.GENERIC_TEXT(msgArray[i],colors[0],bgColor)
	}
	return t
}

func getBorderTile() tile.Tile {
	return tile.GENERIC_TEXT(" ",core.TermCodes(core.FgGreen),core.TermCodes(core.BgGreen))
}

