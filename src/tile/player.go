package tile

import "example/gotell/src/core"

type Player struct {
	UnderTile        Tile
	Tile             Tile
	Name             string
	X, Y, PrvX, PrvY int
}

func GeneratePlayer() Player {
	return Player{
		X:    10,
		Y:    10,
		PrvX: 10,
		PrvY: 10,
		UnderTile: Tile{
			Name:  "ENTRANCE",
			Icon:  core.Icons(core.ICON_LADDER),
			Color: core.TermCodes(core.FgWhite),
		},
		Tile: Tile{
			Name:      "PLAYER",
			Icon:      core.Icons(core.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: core.ATTR_SOLID + core.ATTR_FOREGROUND,
		},
	}
}
