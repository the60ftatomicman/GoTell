package tile

import "example/gotell/src/core"

type Tile struct {
	Name              string         `default:"UNASSIGNED"`
	Icon              core.Icons     `default:Icons(ICON_NULL)`
	Color             core.TermCodes `default:""`
	BGColor           core.TermCodes `default:""`
	Status            string         `default:"OK"`
	Attribute         string         `default:""`
	Parent            *Enemy         `default:nil`
}

var BLANK = Tile{
	Name: "BLANK",
	Icon: core.Icons(core.ICON_BLANK),
	Color:     core.TermCodes(core.BgBlack),
	BGColor:   core.TermCodes(core.BgBlack),
}
var BLANKW = Tile{
	Name: "BLANK",
	Icon: core.Icons(core.ICON_BLANK),
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.FgWhite),
}
var WALL = Tile{
	Name:      "WALL",
	Icon:      core.Icons(core.ICON_WALL),
	Color:     core.TermCodes(core.FgBlue),
	BGColor:   core.TermCodes(core.BgBlue),
	Attribute: core.ATTR_SOLID,
}
var PROFILE_V = Tile{
	Name:      "PROFILE_V",
	Icon:      core.Icons(core.ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var PROFILE_H = Tile{
	Name:      "PROFILE_H",
	Icon:      core.Icons(core.ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var INFO_V = Tile{
	Name:      "INFO_V",
	Icon:      core.Icons(core.ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var INFO_H = Tile{
	Name:      "INFO_H",
	Icon:      core.Icons(core.ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: core.ATTR_SOLID,
}
var LADDER = Tile {
	Name:  "LADDER",
	Icon:  core.Icons(core.ICON_LADDER),
	Color: core.TermCodes(core.FgWhite),
}

func GENERIC_TEXT(character string, colors ...core.TermCodes) Tile{
	bgColor := core.TermCodes(core.BgBlack)
	if(len(colors) > 1){
		bgColor = core.TermCodes(colors[1])
	}
	return Tile{
		Name: "PROFILE_TEXT",
		Icon: core.StringToIcon(character),
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