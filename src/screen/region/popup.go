package region

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strings"
)

func POPUP_LEFT () int { return MAP_LEFT }
func POPUP_TOP () int { return MAP_TOP }
func POPUP_LINES () int { return 1 }
func POPUP_COLUMNS () int { return MAP_COLUMNS }

// Popup
// Used to display a one time menu
type Popup struct {
	Title string
	Message string
	Commands string
	Buffer  [][]tile.Tile
}

func (p *Popup) Initialize(b [][]tile.Tile) {
	//haha yeah I am going to ignore B passed in.
	p.Set("Some dire message!")
	b = p.compile()

	p.Buffer = initializeBuffer(POPUP_LINES(), POPUP_COLUMNS(), b,tile.BLANK)
}

func (p *Popup) Refresh() {
	b := p.compile()
	p.Buffer = initializeBuffer(POPUP_LINES(), POPUP_COLUMNS(), b,tile.BLANK) // rename to generateBuffer?
}

func (p *Popup) Get() (int, int, int, int, [][]tile.Tile) {
	return POPUP_LEFT(), POPUP_TOP(), POPUP_LINES(), POPUP_COLUMNS(), p.Buffer
}

func (p *Popup) Set(msgs string){
	p.Message = msgs
}

func (p *Popup) compile()[][]tile.Tile{
	t := [][]tile.Tile{p.getBaseRow(1,p.Message,core.FgWhite)}
	t = append(t, tile.GenerateHorizontalDivider(POPUP_COLUMNS(),tile.BLANK,tile.INFO_H))
	return t
}

func (p *Popup) getBaseRow(colIdx int, extraMsg string,color core.TermCodes ) []tile.Tile {
	t        := []tile.Tile{} //TODO im lazy and just not making tile
	msgArray := strings.Split(extraMsg, "")
	endIdx   := len(msgArray)+1

	if(endIdx > POPUP_COLUMNS()){
		endIdx = POPUP_COLUMNS()
	}

	for i := 0; i < POPUP_COLUMNS(); i++ {
		if(i >= colIdx && i < endIdx){
			t = append(t, tile.GENERIC_TEXT(msgArray[i-colIdx],color,core.BgBlack))
		}else{
			t = append(t, tile.GENERIC_TEXT(" ",color,core.BgBlack))
		}
	}
	return t
}
