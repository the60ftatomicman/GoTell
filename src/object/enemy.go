package object

import (
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strconv"
	"strings"
)

type Enemy struct {
	Tile             tile.Tile
	Status           string `default:""`
	Name             string `default:""`
	X,Y,XP           int    `default:0`
	Stats            Stats
}

func (e *Enemy) Interaction(s *Stats) bool {
	removeEnemy := false
	// Do dmg
	playerDmg := statCalc_Battle(s,&e.Stats)
	enemyDmg  := statCalc_Battle(&e.Stats,s)
	if(s.GetSpeedWithMod() >= e.Stats.GetSpeedWithMod()){
		e.Stats.UpdateHealth(playerDmg * -1)
		if (e.Stats.Health > 0) {
			s.UpdateHealth(enemyDmg * -1)
		}
	}else{
		s.UpdateHealth(enemyDmg * -1)
		if (s.Health > 0) {
			e.Stats.UpdateHealth(playerDmg * -1)
		}
	}
	//Did player die?
	if (s.Health <= 0) {
		return false
	}
	//Attribute assigning time
	e.applyEffects(s)
	//Test if enemy is dead
	if (e.Stats.Health <= 0) {
		removeEnemy = true
		xpBoost := 0
		if(e.Stats.Level > s.Level){
			xpBoost = e.Stats.Level - s.Level
		}
		s.ChangeXP(e.Stats.XP+xpBoost);
	}
	return removeEnemy
}

func (e *Enemy) applyEffects(s *Stats) {
	if(tile.CheckAttributes(e.Tile,overrides.ATTR_POISONOUS)){
		s.AddEffects(overrides.ATTR_POISONOUS)
	}
	if(tile.CheckAttributes(e.Tile,overrides.ATTR_MANABURN)){
		s.AddEffects(overrides.ATTR_MANABURN)
		s.Mana = 0
	}
} 

func (e *Enemy) CalcDefeat(s *Stats) (int,int) {
	playerDmg := statCalc_Battle(s,&e.Stats)
	enemyDmg  := statCalc_Battle(&e.Stats,s)
	return playerDmg,enemyDmg
}

func (e *Enemy) Convert(s *Stats) {}

func (e *Enemy) GetBufferData() (int,int,string,tile.Tile) {
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
func GenerateEnemiesFromFile(fileData []string) [10][]Enemy {
	//We'll enforce Order of Operatins. enemy 0 is ALWAYS your boss
	enemyList := [10][]Enemy{}
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
     "ENEMY_BOSS_MELEE" : ENEMY_BOSS_MELEE,
     "ENEMY_MELEE"      : ENEMY_MELEE,
	 "ENEMY_POISON"     : ENEMY_POISON,
	 "ENEMY_MANABURN"   : ENEMY_MANABURN,
}

func fileParserEnemy(enemyVals string) ([]string,[]Enemy){
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
func getEnemyTile(t tile.Tile, attrs ...string) tile.Tile{
	for _,attr := range attrs{
		t.Attribute += attr
	}
	return t
}
// Basic Grunt
var ENEMY_MELEE = Enemy{
	Name: "Moleman",
	Stats: Stats{
		MaxHealth: 50,
		HealthMod: 5,
		Health:    50,
		Defense:   5,
		DefMod:    11,
		Offense:   5,
		OffMod:    11,
		Speed:     2,
		FogRet:    30, 
		XP:        5,
	},
	Tile : getEnemyTile(overrides.ENEMY_BASIC),
}

// Boss of the basic grunts
var ENEMY_BOSS_MELEE = Enemy{
	Name: "Boss Moleman",
	Stats: Stats{
		MaxHealth: 50,
		HealthMod: 5,
		Health:    50,
		Defense:   5,
		DefMod:    11,
		Offense:   5,
		OffMod:    11,
		Speed:     2,
		FogRet:    30, 
		XP:        5,
	},
	Tile: getEnemyTile(overrides.ENEMY_BOSS),
}

// Poisonous! TODO, make this a more generic to get the tiles....
var ENEMY_POISON = Enemy{
	Name: "Snake",
	Stats: Stats{
		MaxHealth: 60,
		HealthMod: 5,
		Health:    60,
		Defense:   1,
		DefMod:    5,
		Offense:   15,
		OffMod:    11,
		Speed:     3,
		FogRet:    25, 
		XP:        5,
	},
	Tile : getEnemyTile(overrides.ENEMY_BASIC,overrides.ATTR_POISONOUS),
}

// Mana Burn!
var ENEMY_MANABURN = Enemy{
	Name: "Ghost",
	Stats: Stats{
		MaxHealth: 60,
		HealthMod: 5,
		Health:    60,
		Defense:   1,
		DefMod:    5,
		Offense:   15,
		OffMod:    11,
		Speed:     3,
		FogRet:    25, 
		XP:        5,
	},
	Tile : getEnemyTile(overrides.ENEMY_BASIC,overrides.ATTR_MANABURN),
}