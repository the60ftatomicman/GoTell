package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	"example/gotell/src/object"
)

const MENU_LEFT    = 0
const MENU_TOP     = 0
const MENU_LINES   = 29
const MENU_COLUMNS = 100

// IMPORTANT LINES IN THE MENU FOR WRITING!
// START INDEX 0
// Remember, each of these only have 16 characters!

//const LINE_VAR_NAME       = 1 // Line for player name

// Menu
// Where we are going to setup our character
// This area gives detailed stats about the player
// XP,Level,Health,Mana as well as the items the player currently has
type Menu struct {
	Player *object.Player
	SelectedItem string `default:""`
	Buffer [][]tile.Tile
}

func (p *Menu) Initialize(b [][]tile.Tile) {

	//haha yeah I am going to ignore B passed in.
	b = p.compile()

	p.Buffer = screen.InitializeBuffer(MENU_LINES, MENU_COLUMNS, b,tile.BLANK)
}

func (p *Menu) Get() (int, int, int, int, [][]tile.Tile) {
	return MENU_LEFT, MENU_TOP, MENU_LINES, MENU_COLUMNS, p.Buffer
}

func (p *Menu) ReadDataFromPlayer(plyr *object.Player) [][]tile.Tile {
	p.Player = plyr
	return [][]tile.Tile{}
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

	for i := 0; i < MENU_LINES-3; i++ {
		t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
			tile.GENERIC_TEXT(" ",core.TermCodes(core.FgWhite),core.TermCodes(core.BgBlack)),
			tile.GENERIC_TEXT(" ",core.TermCodes(core.FgWhite),core.TermCodes(core.BgBlack)),
		))
	}

	t = append(t, tile.GenerateHorizontalDivider(MENU_COLUMNS-2,
		getBorderTile(),
		getBorderTile(),
	))

	return t
}
func getBorderTile() tile.Tile{
	return tile.GENERIC_TEXT(" ",core.TermCodes(core.FgGreen),core.TermCodes(core.BgGreen))
}
/*
func (p *Profile)getBaseRow(colIdx int, extraMsg string,colors ...core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{overrides.PROFILE_V}
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
	t = append(t, overrides.PROFILE_V)
	return t
}
*/