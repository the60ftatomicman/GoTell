package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strings"
)

func POPUP_LEFT () int { return MAP_LEFT }
func POPUP_TOP () int { return MAP_TOP }
func POPUP_LINES () int { return 3 }
func POPUP_COLUMNS (msglen int) int { return msglen+4 }

// Popup
// Used to display a one time menu
type Popup struct {
	Title string
	Messages []string //TODO - fix this to be used!
	Commands string
	Buffer  [][]tile.Tile
}
//
//
//
//
//
func (p *Popup) Initialize(b [][]tile.Tile) {
	p.ClearMessages()
	b = p.compile()

	p.Buffer = screen.InitializeBuffer(POPUP_LINES(), POPUP_COLUMNS(0), b,tile.BLANK)
}

func (p *Popup) Refresh() {
	b := p.compile()
	p.Buffer = screen.InitializeBuffer(POPUP_LINES(), POPUP_COLUMNS(len(p.Messages[0])), b,tile.BLANK) // rename to generateBuffer?
}

func (p *Popup) Get() (int, int, int, int, [][]tile.Tile) {
	return POPUP_LEFT(), POPUP_TOP(), POPUP_LINES(), POPUP_COLUMNS(len(p.Messages[0])), p.Buffer
}

func (p *Popup) Set(msgs string){
	p.Messages = []string{}
	p.Messages = append(p.Messages, msgs)
}

func (p *Popup) compile()[][]tile.Tile{
	t := [][]tile.Tile{tile.GenerateHorizontalDivider(POPUP_COLUMNS(len(p.Messages)),tile.BLANK,overrides.INFO_H)}
	if(p.HasMessages()){
		t = append(t,p.getBaseRow(1,p.Messages[0],core.FgWhite))
	}
	t = append(t, tile.GenerateHorizontalDivider(POPUP_COLUMNS(len(p.Messages)),tile.BLANK,overrides.INFO_H))
	return t
}

func (p *Popup) getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{} //TODO im lazy and just not making tile
	msgArray := strings.Split(extraMsg, "")
	endIdx   := len(msgArray)+1

	if(endIdx > POPUP_COLUMNS(len(p.Messages[0]))){
		endIdx = POPUP_COLUMNS(len(p.Messages[0]))
	}

	for i := 0; i < POPUP_COLUMNS(len(p.Messages[0])); i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],color,core.BgBlack))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",color,core.BgBlack))
		}
	}
	return t
}
//
//
//
//
//
func (p *Popup) HasMessages() bool{
	return len(p.Messages) > 0
}
func (p *Popup) ClearMessages() {
	p.Messages = []string{}
}