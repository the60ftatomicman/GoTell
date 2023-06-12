package region

import (
	"bufio"
	"example/gotell/src/core"
	"example/gotell/src/core/tile"
	"os"
	"strconv"
	"strings"
)

const SPLASH_LEFT = 0
const SPLASH_TOP = 0
const SPLASH_LINES = 29
const SPLASH_COLUMNS = 100

// SPLASH
// Used to generate full screen displays like the title screen
type Splash struct {
	FilePath string
	Buffer  [][]tile.Tile
}

func (p *Splash) Initialize(b [][]tile.Tile) {
	p.ParseSplashFile();
}

func (p *Splash) Refresh() {
	p.ParseSplashFile();
}

func (p *Splash) Get() (int, int, int, int, [][]tile.Tile) {
	return SPLASH_LEFT, SPLASH_TOP, SPLASH_LINES, SPLASH_COLUMNS, p.Buffer
}

var colorConverterForeground = map[string]string{
     "w": core.FgWhite,
     "e": core.FgBlack, // -- e for empty I guess :/
	 "b": core.FgBlue,
	 "c": core.FgCyan,
	 "g": core.FgGreen,
	 "m": core.FgMagenta,
	 "r": core.FgRed,
	 "y": core.FgYellow,
	 "d": core.FgGrey, // -- d for dust or dirty
}

var colorConverterBackground = map[string]string{
     "w": core.BgWhite,
     "e": core.BgBlack, // -- e for empty I guess :/
	 "b": core.BgBlue,
	 "c": core.BgCyan,
	 "g": core.BgGreen,
	 "m": core.BgMagenta,
	 "r": core.BgRed,
	 "y": core.BgYellow,
	 "d": core.BgGrey, // -- d for dust or dirty
}

func (p *Splash) ParseSplashFile(){
	readFile,err  := os.Open(p.FilePath)
	if(err != nil){
		print(err)
		panic(err)
	}
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
	
    for fileScanner.Scan() {
		fileLine := fileScanner.Text()
		
		//TODO make this a regex
		if(fileLine == "blank"){
			var blankLine = []tile.Tile{}
			for i := 0;i<SPLASH_COLUMNS;i++ {
				blankLine = append(blankLine, tile.BLANK)
			}
			p.Buffer = append(p.Buffer,blankLine)
		}else{
			var lineData = []tile.Tile{}
			tileStrings := strings.Split(fileLine, ";")
			for _, charStrings := range tileStrings {
				data:= strings.Split(charStrings, ",")
				max,_ := strconv.Atoi(data[0])
				for i:=0;i<max; i++ {
					colors := strings.Split(data[1],"")
					fg := core.TermCodes(colorConverterForeground[colors[0]])
					bg := core.TermCodes(colorConverterBackground[colors[1]])
					for _,letter := range strings.Split(data[2],"") {
						lineData = append(lineData,tile.GENERIC_TEXT(letter,fg,bg))
					}
				}
			}
			p.Buffer = append(p.Buffer,lineData)
		}
    }
    readFile.Close()
}