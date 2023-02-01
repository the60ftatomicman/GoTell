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
	Stats            Stats
}

func (e *Enemy) Interaction(p *Player) bool {
	removeEnemy := false
	if(p.Stats.Speed >= e.Stats.Speed){
		e.Stats.Health -= statCalc_Battle(p.Stats.Offense,e.Stats.Defense)
		p.UpdateHealth(statCalc_Battle(e.Stats.Offense,p.Stats.Defense) * -1)
	}else{
		p.UpdateHealth(statCalc_Battle(e.Stats.Offense,p.Stats.Defense) * -1)
		e.Stats.Health -= statCalc_Battle(p.Stats.Offense,p.Stats.Defense)
	}
	if (e.Stats.Health <= 0) {
		removeEnemy = true
	}
	return removeEnemy
}

func generateEnemy() Enemy {
	e := Enemy{
		Name:      "Moleman",
		X:         10,
		Y:         10,
		PrvX:      10,
		Prvy:      10,
		Stats: Stats{
			Health: 10,
			Defense: 0,
			Offense: 5,
			Speed:   1,
		},
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