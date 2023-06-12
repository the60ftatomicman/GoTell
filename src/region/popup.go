package region

import (
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strings"
)

func POPUP_LEFT () int { return MAP_LEFT + MAP_COLUMNS/3 }
func POPUP_TOP () int { return MAP_TOP + MAP_LINES/3 }
func POPUP_LINES (m []string) int { return len(m)+2 }
func POPUP_COLUMNS (msglen int) int { return msglen+2 }

// Popup
// Used to display a one time menu
type Popup struct {
	Title string
	Messages []string
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

	p.Buffer = screen.InitializeBuffer(POPUP_LINES(p.Messages), POPUP_COLUMNS(0), b,tile.BLANK)
}

func (p *Popup) Refresh() {
	cols := 0
	if(p.HasMessages()){
		cols = p.GetLongestMessage()+1
	}
	b := p.compile()
	p.Buffer = screen.InitializeBuffer(POPUP_LINES(p.Messages), POPUP_COLUMNS(cols), b,tile.BLANK) // rename to generateBuffer?
}

func (p *Popup) Get() (int, int, int, int, [][]tile.Tile) {
	cols := 0
	if(p.HasMessages()){
		cols = p.GetLongestMessage()
	}
	return POPUP_LEFT(), POPUP_TOP(), POPUP_LINES(p.Messages), POPUP_COLUMNS(cols), p.Buffer
}

func (p *Popup) Set(msgs... string){
	p.Messages = []string{}
	for i := 0; i < len(msgs); i++ {
		p.Messages = append(p.Messages, msgs[i])
	}
}

func (p *Popup) compile()[][]tile.Tile{
	closeMessage := "Press y to Dismiss"
	cols := 0
	if(p.HasMessages()){
		p.Messages = append(p.Messages, " ")
		p.Messages = append(p.Messages, closeMessage)
		cols = p.GetLongestMessage()
	}

	t := [][]tile.Tile{tile.GenerateHorizontalDivider(POPUP_COLUMNS(cols)-2,tile.BLANK,overrides.INFO_H)}
	//Ignore the closeMessage
	for i := 0; i < len(p.Messages); i++{
		t = append(t,p.getBaseRow(1,p.Messages[i],core.FgWhite))
	}
	//Add message in divider
	t = append(t, tile.GenerateHorizontalDivider(POPUP_COLUMNS(cols)-2,tile.BLANK,overrides.INFO_H))
	return t
}

func (p *Popup) getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{} //TODO im lazy and just not making tile
	msgArray := strings.Split(extraMsg, "")
	endIdx   := len(msgArray)+1
	cols     := 0
	if(p.HasMessages()){
		cols = p.GetLongestMessage()+1
	}

	if(endIdx > POPUP_COLUMNS(cols)){
		endIdx = POPUP_COLUMNS(cols)
	}

	for i := 0; i < POPUP_COLUMNS(cols); i++ {
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
func (p *Popup) GetLongestMessage() int {
	longest := 0;
	for i := 0; i < len(p.Messages); i++{
		if(longest < len(p.Messages[i])){
			longest = len(p.Messages[i])
		}
	}
	return longest;
}
func (p *Popup) HasMessages() bool{
	return len(p.Messages) > 0
}
func (p *Popup) ClearMessages() {
	p.Messages = []string{}
}