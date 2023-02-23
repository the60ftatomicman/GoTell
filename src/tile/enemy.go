package tile

import (
	"strconv"
	"strings"
)

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

func (e *Enemy) CalcDefeat(s *Stats) int {
	hits := e.Stats.Health / statCalc_Battle(s.Offense,e.Stats.Defense,s.Level)
	if hits < 1 {
		hits = 1
	}
	return hits
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
func GenerateEnemiesFromFile() [10][]Enemy {
	//We'll enforce Order of Operatins. enemy 0 is ALWAYS your boss
	enemyList := [10][]Enemy{}
	fileData := []string{
		"boss:ENEMY_BOSS_MOLEMAN",
		"1,2,3,4,5,6,7,8,9:ENEMY_MOLEMAN",
	}
	for _,row := range fileData {
		indicies,enemies := fileParserEnemy(row)
		
		for _,idx := range indicies{
			idxInt,_ := strconv.Atoi(idx)
			enemyList[idxInt] = append(enemyList[idxInt],enemies...)
		}
	}
	return enemyList
}
//
//
//
//
//
var dataConverterEnemy = map[string]Enemy{
     "ENEMY_BOSS_MOLEMAN": ENEMY_BOSS_MOLEMAN,
     "ENEMY_MOLEMAN": ENEMY_MOLEMAN,
}

func fileParserEnemy(enemyVals string) ([]string,[]Enemy){
	//enemyStrings := strings.Split(enemyVals, ":")
	enemies  := []Enemy{}
	indicies := []string{}
	keyVal   := strings.Split(enemyVals, ":")
	key      := keyVal[0]
	value    := keyVal[1]
	if(key == "boss"){
		indicies = append(indicies, "0")
		enemies  = append(enemies, dataConverterEnemy[value])
	}else{
		indicies = append(indicies, strings.Split(key, ",")...)
		for _,e := range strings.Split(value, ","){
			enemies  = append(enemies, dataConverterEnemy[e])
		}
	}
	return indicies,enemies
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