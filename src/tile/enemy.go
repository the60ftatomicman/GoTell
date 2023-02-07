package tile

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
func (e *Enemy) GetBufferData() (int,int,string,Tile) {
	return e.Y,e.X,e.Name,e.Tile
}
func generateEnemy(x int,y int,l int,e Enemy) Enemy {
	e.X = x
	e.PrvX = x
	e.Y = y
	e.Prvy = y
	e.Stats.Level = l
	return e
}
//
//
//
//
func GenerateEnemiesFromFile() []Enemy{
	return []Enemy{
		generateEnemy(11,11,1,ENEMY_MOLEMAN),
		generateEnemy(11,13,1,ENEMY_MOLEMAN),
		generateEnemy(12,12,4,ENEMY_BOSS_MOLEMAN),
		generateEnemy(13,11,1,ENEMY_MOLEMAN),
		generateEnemy(13,13,1,ENEMY_MOLEMAN),
	}
}
//
//
//
//
var ENEMY_MOLEMAN = Enemy{
	Name:      "Moleman",
	Stats: Stats{
		MaxHealth: 10,
		Health:    10,
		Defense:   0,
		Offense:   5,
		Speed:     1,
		FogRet:    20,
		XP:         5,
	},
	Tile : ENEMY_BASIC,
}
var ENEMY_BOSS_MOLEMAN = Enemy{
	Name:      "Moleman",
	Stats: Stats{
		MaxHealth: 10,
		Health:    10,
		Defense:   0,
		Offense:   5,
		Speed:     1,
		FogRet:    20,
		XP:         5,
	},
	Tile: ENEMY_BOSS,
}