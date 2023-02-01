package core

type Icons string

const (
	ICON_NULL      Icons = "?"
	ICON_PLAYER          = "@"
	ICON_BLANK           = " "
	ICON_WALL            = "#"
	ICON_LADDER          = "H"
	ICON_EMPTY           = "."
	ICON_PROFIlE_V       = "|"
	ICON_PROFIlE_H       = "-"
	ICON_FOG             = "?"
)

func StringToIcon(character string) Icons {
	return Icons(character)
}
