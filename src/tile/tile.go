package tile

import "example/gotell/src/core"

type Tile struct {
	Name              string         `default:"UNASSIGNED"`
	Icon              core.Icons     `default:Icons(ICON_NULL)`
	Color             core.TermCodes `default:""`
	Status            string         `default:"OK"`
	Attribute, Object string         `default:""`
	Parent            *Enemy         `default:nil`
}

var BLANK = Tile{
	Name: "BLANK",
	Icon: core.Icons(core.ICON_BLANK),
}
var WALL = Tile{
	Name:      "WALL",
	Icon:      core.Icons(core.ICON_WALL),
	Color:     core.TermCodes(core.FgBlue),
	Attribute: core.ATTR_SOLID,
}
var PROFILE_V = Tile{
	Name:      "PROFILE_Vs",
	Icon:      core.Icons(core.ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	Attribute: core.ATTR_SOLID,
}
