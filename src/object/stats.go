package object

import (
	overrides "example/gotell/src/core_overrides"
	"strings"
)

type Stats struct {
	MaxHealth int `default:100` // Maximum health the entity can have
	Health    int `default:100` // Current health the entity has
	MaxMana   int `default:100` // Maximum mana the entity can have
	Mana      int `default:100` // Current mana the entity has
	Level     int `default:1`   // Current level the entity has
	LevelMod  int `default:10`  // Amount of XP required for entity to reach the next level
	XP        int `default:0`   // Current XP
	Defense   int `default:1`   // Used in calculating how much dmg an entity takes
	Offense   int `default:1`   // Use in calculating how much dmg an entity gives
	Speed     int `default:1`   // Determines turn order in combat
	Favor     int `default:0`   // Like mana for Gods. TODO -- implement using this!
	FogRet    int `default:25`  // How much HEALTH and MANA is returned when a FOG tile is removed
	Vision    int `default:1`   // How FAR the entity can see, used in FOG removal
	Effects   string `default:"'` // Like attributes
}

//TODO -- at the moment this "works" but my enemies and player need to have their stats modified
// to really make this AOK
func statCalc_Battle_Percentage(off int, def int, offMod int,defHP int) int{
	dmg := 0
	if((off * offMod) > def) {
		var offPercentage float64   = (float64(off) * float64(offMod)) / 100.0
		var defPercentage float64   = float64(def) / 100.0
		var deltaPercentage float64 = offPercentage - defPercentage
		var healthPoints    float64 = float64(defHP) * deltaPercentage
		if healthPoints < 1.0 {healthPoints = 1}
		dmg = int(healthPoints)
	}
	return dmg
}

func statCalc_Battle(off int, def int, offMod int) int{
	dmg := 0
	if((off * offMod) > def){
		dmg = (off * offMod) - def
	}
	return dmg
}

//convert to %
func (s *Stats) UpdateHealth(delta int) {
	if s.checkEffects(overrides.ATTR_POISONOUS) && delta > 0 {
		delta = 0
	}
	s.Health += delta
	if (s.Health > s.MaxHealth) {
		s.Health = s.MaxHealth
	}
	if (s.Health < 0) {
		s.Health = 0
	}
}

//convert to %
func (s *Stats) UpdateMana(delta int) {
	if s.checkEffects(overrides.ATTR_MANABURN) && delta > 0 {
		delta = 0
	}
	s.Mana += delta
	if (s.Mana > s.MaxMana) {
		s.Mana = s.MaxMana
	}
	if (s.Mana < 0) {
		s.Mana = 0
	}
}

func (s *Stats)ChangeXP(deltaXP int) {
	s.XP += deltaXP
	if(s.XP >= s.LevelMod) {
		s.Level  = (s.XP /  s.LevelMod) + s.Level
		s.XP     = s.XP %  s.LevelMod
		s.Health = s.MaxHealth
		s.Mana   = s.MaxMana
	}
}

func (s *Stats)checkEffects(attr string) bool{
	return strings.Contains(s.Effects, attr)
}
func (s *Stats)AddEffects(attrs ...string){
	for _,attr := range attrs{
		if(!s.checkEffects(attr)){
			s.Effects += attr
		}
	}
}
func (s *Stats)RemoveEffects(attrs ...string){
	for _,attr := range attrs{
		s.Effects = strings.Replace(s.Effects,attr,"",1)
	}
}