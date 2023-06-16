package tile

import (
	"example/gotell/src/core"
	"strings"
)

type Tile struct {
	Name              string              `default:"UNASSIGNED"`
	Icon              Icons               `default:Icons(ICON_NULL)`
	Color             core.TermCodes      `default:""`
	BGColor           core.TermCodes      `default:""`
	Attribute         string              `default:""`
}
//
//
//
//
//
var BLANK = Tile{
	Name:      "BLANK",
	Icon: 	   Icons(ICON_BLANK),
	Color:     core.TermCodes(core.BgBlack),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: "",
}
var NULL = Tile{
	Name:      "NULL",
	Icon:      Icons(ICON_NULL),
	Color:     core.TermCodes(core.FgYellow),
	BGColor:   core.TermCodes(core.BgMagenta),
	Attribute: "",
}
//
//
//
//
//
func GENERIC_TEXT(character string, colors ...core.TermCodes) Tile{
	bgColor := core.TermCodes(core.BgBlack)
	if(len(colors) > 1){
		bgColor = core.TermCodes(colors[1])
	}
	return Tile{
		Name: "GENERIC_TEXT",
		Icon: StringToIcon(character),
		Color: core.TermCodes(colors[0]),
		BGColor: bgColor,
	}
}

func GenerateHorizontalDivider(length int,bookend Tile,fill Tile) []Tile {
	t := []Tile{bookend}
	for i := 0; i < length; i++ {
		t = append(t, fill)
	}
	t = append(t, bookend)
	return t
}

func CheckAttributes(t Tile, attr Attributes) bool{
	return strings.Contains(t.Attribute, string(Attributes(attr)))
}
