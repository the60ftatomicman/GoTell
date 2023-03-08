package overrides

import (
	"example/gotell/src/core"
	"example/gotell/src/core/tile"
)

var BLANK = tile.Tile{
	Name: "BLANK",
	Icon: 	   tile.Icons(tile.ICON_BLANK),
	Color:     core.TermCodes(core.BgBlack),
	BGColor:   core.TermCodes(core.BgBlack),
}
var NULL = tile.Tile{
	Name: "NULL",
	Icon: tile.Icons(tile.ICON_NULL),
	Color:     core.TermCodes(core.FgYellow),
	BGColor:   core.TermCodes(core.BgMagenta),
}
var WALL = tile.Tile{
	Name:      "WALL",
	Icon:      tile.Icons(ICON_WALL),
	Color:     core.TermCodes(core.FgBlue),
	BGColor:   core.TermCodes(core.BgBlue),
	Attribute: ATTR_SOLID,
}
var PROFILE_V = tile.Tile{
	Name:      "PROFILE_V",
	Icon:      tile.Icons(ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_SOLID,
}
var PROFILE_H = tile.Tile{
	Name:      "PROFILE_H",
	Icon:      tile.Icons(ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_SOLID,
}
var INFO_V = tile.Tile{
	Name:      "INFO_V",
	Icon:      tile.Icons(ICON_PROFIlE_V),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_SOLID,
}
var INFO_H = tile.Tile{
	Name:      "INFO_H",
	Icon:      tile.Icons(ICON_PROFIlE_H),
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_SOLID,
}
//
var LADDER = tile.Tile{
	Name:  "LADDER",
	Icon:  tile.Icons(ICON_LADDER),
	Color: core.TermCodes(core.FgYellow),
	Attribute: ATTR_CLIMBABLE,
}
var EMPTY = tile.Tile{
	Name:  "EMPTY",
	Icon:  tile.Icons(ICON_EMPTY),
	Color: core.TermCodes(core.FgWhite),
}
var FOG = tile.Tile{
	Name:  "FOG",
	Icon:  tile.Icons(ICON_FOG),
	Color: core.TermCodes(core.FgGrey),
	BGColor: core.TermCodes(core.BgGrey),
	Attribute: ATTR_FOREGROUND,
}
//
var ENEMY_BASIC = tile.Tile{
	Name:      "ENEMY",
	Icon:      ICON_ENEMY,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: ATTR_FIGHTABLE + ATTR_SOLID,
}
var ENEMY_BOSS = tile.Tile{
	Name:      "ENEMY",
	Icon:      ICON_BOSS,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: ATTR_FIGHTABLE + ATTR_SOLID + ATTR_BOSS,
}
var ENEMY_SPAWN = tile.Tile{
	Name:      "ENEMY_SPAWN",
	Icon:      tile.ICON_NULL,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: "",
}

var ITEM_SPAWN = tile.Tile{
	Name:      "ITEM_SPAWN",
	Icon:      tile.ICON_NULL,
	Color:     core.TermCodes(core.FgWhite),
	BGColor:   core.TermCodes(core.BgRed),
	Attribute: "",
}

var POTION_HEALTH = tile.Tile{
	Name:      "ITEM",
	Icon:      ICON_HEALTH,
	Color:     core.TermCodes(core.FgRed),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_USABLE,
}
var POTION_MANA = tile.Tile{
	Name:      "ITEM",
	Icon:      ICON_MANA,
	Color:     core.TermCodes(core.FgCyan),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_USABLE,
}
var EQUIPTMENT= tile.Tile{
	Name:      "ITEM",
	Icon:      ICON_EQUIPTMENT,
	Color:     core.TermCodes(core.FgYellow),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_EQUIPTABLE,
}
var SPELL = tile.Tile{
	Name:      "ITEM",
	Icon:      ICON_SPELL,
	Color:     core.TermCodes(core.FgMagenta),
	BGColor:   core.TermCodes(core.BgBlack),
	Attribute: ATTR_USABLE+ATTR_SPELL,
}