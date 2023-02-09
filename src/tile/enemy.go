package tile

type Enemy struct {
	Tile             Tile
	Status           string `default:""`
	Name             string `default:""`
	X,Y,XP           int    `default:0`
	Stats            Stats
}

func (e *Enemy) Interaction(s *Stats) bool {
	removeEnemy := false
	if(s.Speed >= e.Stats.Speed){
		e.Stats.UpdateHealth(statCalc_Battle(s.Offense,e.Stats.Defense,s.Level) * -1)
		s.UpdateHealth(statCalc_Battle(e.Stats.Offense,s.Defense,e.Stats.Level) * -1)
	}else{
		s.UpdateHealth(statCalc_Battle(e.Stats.Offense,s.Defense,e.Stats.Level) * -1)
		e.Stats.UpdateHealth(statCalc_Battle(s.Offense,s.Defense,s.Level) * -1)
	}
	if (e.Stats.Health <= 0) {
		removeEnemy = true
		s.ChangeXP(e.Stats.XP);
	}
	return removeEnemy
}

func (e *Enemy) Convert(s *Stats) {}
func (e *Enemy) GetBufferData() (int,int,string,Tile) {
	return e.Y,e.X,e.Name,e.Tile
}
func generateEnemy(x int,y int,l int,e Enemy) Enemy {
	e.X = x
	e.Y = y
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
	Name:      "Boss Moleman",
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