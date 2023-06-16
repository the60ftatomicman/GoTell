package tile

// Icons
// These are used to determine what a tile SHOWS
// *Most* tiles use these overrides. Messages use the function StringToIcons
type Icons string

const (
	ICON_NULL  Icons = "?"
	ICON_BLANK       = " "
)

func StringToIcon(character string) Icons {
	return Icons(character)
}