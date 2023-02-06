package tile

import "example/gotell/src/core"

type Enemy struct {
	Tile             Tile
	Status           string
	Name             string
	X, Y, PrvX, Prvy,XP int
	Stats            Stats
}

func (e *Enemy) Interaction(p *Player) bool {
	removeEnemy := false
	if(p.Stats.Speed >= e.Stats.Speed){
		e.Stats.UpdateHealth(statCalc_Battle(p.Stats.Offense,e.Stats.Defense,p.Stats.Level) * -1)
		p.Stats.UpdateHealth(statCalc_Battle(e.Stats.Offense,p.Stats.Defense,e.Stats.Level) * -1)
	}else{
		p.Stats.UpdateHealth(statCalc_Battle(e.Stats.Offense,p.Stats.Defense,e.Stats.Level) * -1)
		e.Stats.UpdateHealth(statCalc_Battle(p.Stats.Offense,p.Stats.Defense,p.Stats.Level) * -1)
	}
	if (e.Stats.Health <= 0) {
		removeEnemy = true
		p.ChangeXP(e.Stats.XP);
	}
	return removeEnemy
}

func (e *Enemy) Convert(p *Player) {}

func generateEnemy(x int,y int,level int,name string) Enemy {
	e := Enemy{
		Name:      name,
		X:         x,
		Y:         y,
		PrvX:      x,
		Prvy:      y,
		Stats: Stats{
			Level:     level,
			MaxHealth: 10,
			Health:    10,
			Defense:   0,
			Offense:   5,
			Speed:     1,
			FogRet:    20,
			XP:         5,
		},
	}
	e.Tile = Tile{
		Name:      "ENEMY",
		Icon:      "E",
		Color:     core.TermCodes(core.FgRed),
		Attribute: core.ATTR_FIGHTABLE + core.ATTR_SOLID + core.ATTR_FOREGROUND,
	//	Parent:    &e,
	}
	return e
}
//
//
//
//
func GenerateEnemiesFromFile() []Enemy{
	return []Enemy{
		generateEnemy(11,11,1,"Moleman"),
		generateEnemy(11,13,1,"Moleman"),
		generateEnemy(12,12,4,"Moleman Boss"),
		generateEnemy(13,11,1,"Moleman"),
		generateEnemy(13,13,1,"Moleman"),
	}
}