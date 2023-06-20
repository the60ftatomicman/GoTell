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

// IMPORTANT LINES IN THE MENU FOR WRITING!
// START INDEX 0
// Remember, each of these only have 16 characters!

const MENU_LINE_VAR_CLASS = 1 // Line for selecting a class
const MENU_LINE_VAR_LEVEL = 3 // Line for selecting the level
var MENU_OPTIONS_CLASS  = []object.PlayerClass{
	object.CLASS_PHYSICAL,
	object.CLASS_MAGIC,
	object.CLASS_SPEED,
}
var MENU_OPTIONS_LEVEL  = []string{"demolevel","Level B","Level C"} //TODO -- obviously this is getting turned into a struct! 

// Menu
// Where we are going to setup our character
// This area gives details on what the player is about to suggest.
type Menu struct {
	Player *object.Player
	CursorIdx int `default:0`
	CursorClass int `default:0`
	CursorLevel int `default:0`
	Buffer [][]tile.Tile
}

func (p *Menu) Initialize(b [][]tile.Tile) {

	//haha yeah I am going to ignore B passed in.
	b = p.compile()
	p.CursorIdx = 0
	p.Buffer = screen.InitializeBuffer(MENU_LINES, MENU_COLUMNS, b,tile.BLANK)
}

func (p *Menu) Get() (int, int, int, int, [][]tile.Tile) {
	return MENU_LEFT, MENU_TOP, MENU_LINES, MENU_COLUMNS, p.Buffer
}

func (p *Menu) MoveCursor(delta int) {
	p.CursorIdx += delta
	menuItems := 1 // total minus 1
	if (p.CursorIdx < 0){
		p.CursorIdx = menuItems
	} else if (p.CursorIdx > menuItems){
		p.CursorIdx = 0
	}
}

func (p *Menu) ChangeClass(delta int) {
	p.CursorClass += delta
	menuItems := 2 // total minus 1
	if (p.CursorClass < 0){
		p.CursorClass = menuItems
	} else if (p.CursorClass > menuItems){
		p.CursorClass = 0
	}
}

func (p *Menu) ChangeLevel(delta int) {
	p.CursorLevel += delta
	menuItems := 2 // total minus 1
	if (p.CursorLevel < 0){
		p.CursorLevel = menuItems
	} else if (p.CursorLevel > menuItems){
		p.CursorLevel = 0
	}
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

	t = append(t, p.getBaseRow(1,"----- Choose your settings ----- ",core.FgWhite))

	for i := 0; i < MENU_LINES-4; i++ {
		switch(i){
			case MENU_LINE_VAR_CLASS:{
				fgColor := core.FgWhite
				bgColor := core.BgBlack
				if(p.CursorIdx == 0){
					fgColor = core.FgBlue
					bgColor = core.BgWhite		
				}
				t = append(t, p.getBaseRow(1," CLASS: "+MENU_OPTIONS_CLASS[p.CursorClass].Name,core.TermCodes(fgColor),core.TermCodes(bgColor)))
				p.Player.Class = MENU_OPTIONS_CLASS[p.CursorClass].Name
				p.Player.Stats = MENU_OPTIONS_CLASS[p.CursorClass].Stats
			}
			case MENU_LINE_VAR_LEVEL:{
				fgColor := core.FgWhite
				bgColor := core.BgBlack
				if(p.CursorIdx == 1){
					fgColor = core.FgBlue
					bgColor = core.BgWhite		
				}
				t = append(t, p.getBaseRow(1," LEVEL: "+MENU_OPTIONS_LEVEL[p.CursorLevel],core.TermCodes(fgColor),core.TermCodes(bgColor)))
			}
			default: {
				t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
					tile.GENERIC_TEXT(" ",core.TermCodes(core.FgWhite),core.TermCodes(core.BgBlack)),
					tile.GENERIC_TEXT(" ",core.TermCodes(core.FgWhite),core.TermCodes(core.BgBlack)),
				))
			}
		}

	}

	t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
		getBorderTile(),
		getBorderTile(),
	))

	return t
}

func (p *Menu)getBaseRow(colIdx int, extraMsg string,colors ...core.TermCodes ) []tile.Tile {
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

	for i := 0; i < MENU_COLUMNS-1; i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],colors[0],bgColor))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",colors[0],core.BgBlack))
		}
	}
	return t

}

func getBorderTile() tile.Tile {
	return tile.GENERIC_TEXT(" ",core.TermCodes(core.FgGreen),core.TermCodes(core.BgGreen))
}

