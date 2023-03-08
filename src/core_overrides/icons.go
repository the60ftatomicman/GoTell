package overrides

import "example/gotell/src/core/tile"

//These are icons specific to my current game.
const (
	//Non map regions
	ICON_PROFIlE_V tile.Icons  = "|"
	ICON_PROFIlE_H = "-"
	//MAP
	ICON_WALL   = "#"
	ICON_LADDER = "H"
	ICON_FOG    = "?"
	ICON_EMPTY  = "."
	ICON_PLAYER = "@"
	//Enemies
	ICON_ENEMY = "E"
	ICON_BOSS  = "B"
	//Items
	ICON_MANA       = "M"
	ICON_HEALTH     = "H"
	ICON_EQUIPTMENT = "I"
	ICON_SPELL      = "S"
)