package core

import (
	"net"
	"strconv"
	"strings"
)

type TermCodes string

// / https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
const (
	FgBlack     TermCodes = "\u001b[30m"
	FgBlue                = "\u001b[34m"
	FgCyan                = "\u001b[36m"
	FgGreen               = "\u001b[32m"
	FgMagenta             = "\u001b[35m"
	FgRed                 = "\u001b[31m"
	FgWhite               = "\u001b[37m"
	FgReset               = "\u001b[0m"
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

func Clear() string {
	return string(TermCodes(ClearScreen))
}

func GenChar(msg string, color TermCodes) string {
	clr := string(color)
	rst := string(TermCodes(FgReset))

	return clr + msg + rst
}

//func genLine(msg string) string {
//	return msg + string(TermCodes(Newline))
//}

// tile.MAP_LINES+1
func HandleOutputToClient(c net.Conn, cursorX int, cursorY int, out string) {
	c.Write([]byte(Clear() + out + moveCursor(cursorX, cursorY) + ">"))
}
