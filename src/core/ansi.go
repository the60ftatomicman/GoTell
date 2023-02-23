package core

import (
	"net"
	"strconv"
	"strings"
)

// TermCodes
// These are used to generate our colorful text in the terminal
// Most of my macro stuff is pulled from the article belows
// https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
type TermCodes string
const (
	ColorCode   TermCodes =  "\u001b[<FG><BG>m"
	FgReset               = "\u001b[0m"
	//
	FgBlack               = "30"
	FgBlue                = "34"
	FgCyan                = "36"
	FgGreen               = "32"
	FgMagenta             = "35"
	FgRed                 = "31"
	FgWhite               = "37"
	FgYellow              = "33"
	FgGrey                = "90"
	//
	BgBlack               = ";40"
	BgBlue                = ";44"
	BgCyan                = ";46"
	BgGreen               = ";42"
	BgMagenta             = ";45"
	BgRed                 = ";41"
	BgWhite               = ";47"
	BgYellow              = ";43"
	BgGrey                = ";100"
	//
	Newline               = "\r\n"
	ClearScreen           = "\033\143"
	Cursorf               = "\033[<L>;<C>f" // DOES NOTHING ATM
	//CursorH               = "\033[<L>;<C>H" // DOES NOTHING ATM
	CursorRight = "\033[<N>C"
	CursorLeft  = "\033[<N>D"
	CursorDown  = "\033[<N>B"
	CursorUp    = "\033[<N>A"
	ScreenSize  = "\033[8;<L>;<C>t"
	//CursorSave              = "\033[s" // DOES NOTHING ATM
	//CursorRestore           = "\033[u" // DOES NOTHING ATM
)

// ResizeTerminal 
// Changes the window (on users desktop) to given size
// lines == Y or HEIGHT ; columns = X or WIDTH
func ResizeTerminal(lines int, columns int) string {
	template := "<L>;<C>"
	mod := strconv.Itoa(lines) + ";" + strconv.Itoa(columns)
	size := strings.Replace(string(TermCodes(ScreenSize)), template, mod, 1)
	return size
}

// I'll have to use the coordinate version
func moveCursor(x int, y int) string {
	template := "<L>;<C>"
	mod := strconv.Itoa(y) + ";" + strconv.Itoa(x)
	pos := strings.Replace(string(TermCodes(Cursorf)), template, mod, 1)
	return pos
}

// Clear
// Clears the screen of all characters
func Clear() string {
	return string(TermCodes(ClearScreen))
}

// GenChar
// Given a string, apply the Terminal Codes to it
// If no colors are provided, it uses default black background white foreground
// If 1 color is provided we only set the Foreground
func GenChar(msg string, colors ...TermCodes) string {
	fgclr := string(colors[0])
	bgclr := string(TermCodes(BgBlack))
	if(len(colors) > 1){
		bgclr = string(colors[1])
	}
	rst := string(TermCodes(FgReset))
	color := strings.Replace(strings.Replace(string(TermCodes(ColorCode)), "<FG>",fgclr,1),"<BG>",bgclr,1)
	return color + msg + rst
}

// HandleOutputToClient
// Clear the screen and draw the proper console character.
// Note: This does NOT handle input parsing, that is done in the session
func HandleOutputToClient(c net.Conn, cursorX int, cursorY int, out string) {
	c.Write([]byte(Clear() + out + moveCursor(cursorX, cursorY) + ">"))
}
