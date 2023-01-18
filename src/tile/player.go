package tile

import "example/gotell/src/core"

type Player struct {
	UnderTile        Tile
	Tile             Tile
	Name             string
	X, Y, PrvX, PrvY int
	Stats            Stats
}

func GeneratePlayer() Player {
	return Player{
		X:    1,
		Y:    5,
		PrvX: 1,
		PrvY: 5,
		Stats: Stats{
			Health: 100,
			Defense: 1,
			Offense: 1,
			Speed:   2,
		},
		Tile: Tile{
			Name:      "PLAYER",
			Icon:      core.Icons(core.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: core.ATTR_SOLID + core.ATTR_FOREGROUND,
		},
	}
}
