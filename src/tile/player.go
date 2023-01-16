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
		X:    1,
		Y:    5,
		PrvX: 1,
		PrvY: 5,
		Tile: Tile{
			Name:      "PLAYER",
			Icon:      core.Icons(core.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: core.ATTR_SOLID + core.ATTR_FOREGROUND,
		},
	}
}
