package core

type Icons string

const (
	ICON_NULL  Icons = "?"
	ICON_BLANK       = " "
	//Non map regions
	ICON_PROFIlE_V = "|"
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

func StringToIcon(character string) Icons {
	return Icons(character)
}
