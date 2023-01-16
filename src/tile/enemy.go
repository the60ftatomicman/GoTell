package tile

import "example/gotell/src/core"

type iInteractiveObject interface {
	Interaction() Tile
}

type Enemy struct {
	Tile             Tile
	Status           string
	Name             string
	X, Y, PrvX, Prvy int
}

func (e *Enemy) Interaction() bool {
	return true
}

func generateEnemy() Enemy {
	e := Enemy{
		X:         20,
		Y:         12,
		PrvX:      12,
		Prvy:      12,
	}
	e.Tile = Tile{
		Name:      "ENEMY",
		Icon:      "E",
		Color:     core.TermCodes(core.FgRed),
		Attribute: core.ATTR_FIGHTABLE + core.ATTR_SOLID + core.ATTR_FOREGROUND,
		Parent:    &e,
	}
	return e
}
//
//
//
//
func GenerateEnemiesFromFile() []Enemy{
	return []Enemy{generateEnemy()}
}