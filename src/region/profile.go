package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"example/gotell/src/object"
	"strconv"
	"strings"
)

const PROFILE_LEFT    = 80
const PROFILE_TOP     = 0
const PROFILE_LINES   = 29
const PROFILE_COLUMNS = 18

// IMPORTANT LINES IN THE PROFILE FOR WRITING!
// START INDEX 0
// Remember, each of these only have 16 characters!

const LINE_VAR_NAME       = 1 // Line for player name
const LINE_VAR_CLASS      = 2 // Line for player class
const LINE_VAR_HEALTH     = 4 // Line for player health
const LINE_VAR_MANA       = 5 // Line for player mana
const LINE_VAR_OFFENSE    = 7  // Line for player OFFENSE stat
const LINE_VAR_DEFENSE    = 8  // Line for player DEFENSE stat
const LINE_VAR_LEVEL      = 10 // Line for player DEFENSE stat
const LINE_VAR_XP         = 11 // Line for player XP stat
const LINE_LBL_AILMENTS   = 13  // line for label to denote where ailments begin 
const LINE_VAR_AILMENTS   = 14  // line for label to denote where ailments begin
const LINE_VAR_AILMENT_COUNT = 2  // How many item lines we'll display
const LINE_LBL_ITEMS      = 17 // Line for label to denote where items begin
const LINE_VAR_ITEM       = 18 // Starting line for items
var LINE_VAR_ITEM_COUNT   = func(p *object.Player)int{return p.Stats.ItemSlots} // How many item lines we'll display


// Profile
// Displays to the RIGHT of the level
// This area gives detailed stats about the player
// XP,Level,Health,Mana as well as the items the player currently has
type Profile struct {
	Player *object.Player
	SelectedItem string `default:""`
	Buffer [][]tile.Tile
}

func (p *Profile) Initialize(b [][]tile.Tile) {

	//haha yeah I am going to ignore B passed in.
	b = p.compile()

	p.Buffer = screen.InitializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, b,tile.BLANK)
}

func (p *Profile) Get() (int, int, int, int, [][]tile.Tile) {
	return PROFILE_LEFT, PROFILE_TOP, PROFILE_LINES, PROFILE_COLUMNS, p.Buffer
}

func (p *Profile) ReadDataFromPlayer(plyr *object.Player) [][]tile.Tile {
	p.Player = plyr
	return [][]tile.Tile{}
}

func (p *Profile)Refresh(){
	p.Buffer = screen.InitializeBuffer(PROFILE_LINES, PROFILE_COLUMNS, p.compile(),tile.BLANK)
}

func (p *Profile)compile()[][]tile.Tile{
	t := [][]tile.Tile{{tile.BLANK}}
	t = append(t, tile.GenerateHorizontalDivider(PROFILE_COLUMNS-2,tile.BLANK,overrides.PROFILE_H))
	for i := 0; i < PROFILE_LINES-3; i++ {
		switch(i){
			case LINE_VAR_NAME:{
				t = append(t, p.getBaseRow(1," NAME: "+p.Player.Name,core.FgWhite))
			}
			case LINE_VAR_CLASS:{
				t = append(t, p.getBaseRow(1,"CLASS: "+p.Player.Class,core.FgWhite))
			}
			case LINE_VAR_HEALTH:{
				t = append(t, p.getBaseRow(1,"   HP: "+strconv.Itoa(p.Player.Stats.Health)+"/"+strconv.Itoa(p.Player.Stats.GetHealthWithMod()),core.FgRed))
			}
			case LINE_VAR_MANA:{
				t = append(t, p.getBaseRow(1," MANA: "+strconv.Itoa(p.Player.Stats.Mana)+"/"+strconv.Itoa(p.Player.Stats.GetManaWithMod()),core.FgBlue))
			}
			case LINE_VAR_LEVEL:{
				t = append(t, p.getBaseRow(1,"LEVEL: "+strconv.Itoa(p.Player.Stats.Level),core.FgYellow))
			}
			case LINE_VAR_XP:{
				t = append(t, p.getBaseRow(1,"   XP: "+strconv.Itoa(p.Player.Stats.XP)+"/"+strconv.Itoa(p.Player.Stats.LevelMod),core.FgYellow))
			}
			case LINE_VAR_OFFENSE:{
				t = append(t, p.getBaseRow(1,"  OFF: "+strconv.Itoa(p.Player.Stats.GetOffenseWithMod())+" ["+strconv.Itoa(p.Player.Stats.OffItemMod)+"]",core.FgWhite))
			}
			case LINE_VAR_DEFENSE:{
				t = append(t, p.getBaseRow(1,"  DEF: "+strconv.Itoa(p.Player.Stats.GetDefenseWithMod())+" ["+strconv.Itoa(p.Player.Stats.DefItemMod)+"]",core.FgWhite))
			}
			//case LINE_VAR_GOLD:{
			//	t = append(t, p.getBaseRow(1,"GOLD: "+p.Gold,core.FgYellow))
			//}
			case LINE_LBL_ITEMS:{
				slotsLeft := strconv.Itoa(LINE_VAR_ITEM_COUNT(p.Player) - len(p.Player.Items))
				t = append(t, p.getBaseRow(0," - ITEMS ["+slotsLeft+"] - ",core.FgCyan))
			}
			case LINE_LBL_AILMENTS:{
				t = append(t, p.getBaseRow(0," - AILMENTS - ",core.FgMagenta))
			}
			case LINE_VAR_AILMENTS: {
				if(p.Player.Stats.IsManaBurned()){
					t = append(t, p.getBaseRow(0," o- ManaBurned ",core.FgBlue))
				}
			}
			case LINE_VAR_AILMENTS+1: {
				if(p.Player.Stats.IsPoisoned()){
					t = append(t, p.getBaseRow(0," o- Poisoned ",core.FgGreen))
				}
			}
			default:{
				//Assuume bottom is for items.
				item_idx := (i - LINE_VAR_ITEM)
				item_idx_str := strconv.Itoa(item_idx+1)
				if (i >= LINE_VAR_ITEM && i < LINE_VAR_ITEM+LINE_VAR_ITEM_COUNT(p.Player) && item_idx < len(p.Player.Items)){
					if(p.SelectedItem == item_idx_str){
						t = append(t, p.getBaseRow(1,item_idx_str+") "+p.Player.Items[item_idx].Name,core.FgWhite,core.BgBlue))
					}else{
						t = append(t, p.getBaseRow(1,item_idx_str+") "+p.Player.Items[item_idx].Name,core.FgBlack))
					}
				}else{
					// Just append nothing.
					t = append(t, p.getBaseRow(0,"",core.FgBlue))
				}
			}
		}
		
	}
	t = append(t, tile.GenerateHorizontalDivider(PROFILE_COLUMNS-2,tile.BLANK,overrides.PROFILE_H))
	return t
}

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

