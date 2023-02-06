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
	//Items
	ICON_MANA      = "M"
	ICON_HEALTH    = "H"
	ICON_OFF_BOOST = "I"
)

func StringToIcon(character string) Icons {
	return Icons(character)
}
