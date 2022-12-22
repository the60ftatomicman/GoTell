package tile

import "example/gotell/src/core"

type iEnemy interface {
	Interaction() Tile
}

type Enemy struct {
	UnderTile        Tile
	Tile             Tile
	Status           string
	Name             string
	X, Y, PrvX, Prvy int
}

func (e *Enemy) Interaction() bool {
	e.Tile = e.UnderTile
	return true
}

func GenerateEnemy() Enemy {
	e := Enemy{
		X:         12,
		Y:         12,
		PrvX:      12,
		Prvy:      12,
		UnderTile: BLANK,
	}
	e.Tile = Tile{
		Name:      "ENEMY",
		Icon:      "E",
		Color:     core.TermCodes(core.FgGreen),
		Attribute: core.ATTR_FIGHTABLE + core.ATTR_SOLID,
		Parent:    &e,
	}
	return e
}
