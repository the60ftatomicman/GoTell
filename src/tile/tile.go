package tile

import "example/gotell/src/core"

type Tile struct {
	Name              string         `default:"UNASSIGNED"`
	Icon              core.Icons     `default:Icons(ICON_NULL)`
	Color             core.TermCodes `default:""`
	BGColor           core.TermCodes `default:""`
	Status            string         `default:"OK"`
	Attribute, Object string         `default:""`
	Parent            *Enemy         `default:nil`
}

var BLANK = Tile{
	Name: "BLANK",
	Icon: core.Icons(core.ICON_BLANK),
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgBlack),
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


func PROFILE_TEXT(character string, color core.TermCodes) Tile{
	return Tile{
		Name: "PROFILE_TEXT",
		Icon: core.StringToIcon(character),
		Color: core.TermCodes(color),
		BGColor: core.TermCodes(core.BgBlack),
	}
}