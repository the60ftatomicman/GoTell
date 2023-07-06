package object

import (
	overrides "example/gotell/src/core_overrides"
	"strings"
)

type Stats struct {
	MaxHealth     int `default:1`   // Maximum health the entity can have
	Health        int `default:1`   // Current health the entity has
	HealthMod     int `default:0`   // Current health increase per level
	HealthItemMod int `default:0`   // Mod to HP from items
	MaxMana       int `default:0`   // Maximum mana the entity can have
	ManaMod       int `default:0`   // Mana increase per level
	Mana          int `default:0`   // Current mana the entity has
	ManaItemMod   int `default:0`   // Mod to Mana from items
	Level         int `default:1`   // Current level the entity has
	LevelMod      int `default:1`   // Amount of XP required for entity to reach the next level
	XP            int `default:0`   // Current XP
	Defense       int `default:1`   // Used in calculating how much dmg an entity takes
	DefMod        int `default:0`   // Is multiplied by level than added to Defense.
	DefItemMod    int `default:0`   // Modified Defense value added from items
	Offense       int `default:1`   // Use in calculating how much dmg an entity gives
	OffMod        int `default:0`   // Is multiplied by level than added to Offense.
	OffItemMod    int `default:0`   // Modified Offense value added from items
	Speed         int `default:1`   // Determines turn order in combat
	SpeedMod      int `default:0`   // How much speed increase we get per level!
	SpeedItemMod  int `default:0`   // How much speed increase we get based on items.
	Favor         int `default:0`   // Like mana for Gods. TODO -- implement using this!
	FogRet        int `default:1`   // How much HEALTH and MANA is returned when a FOG tile is removed
	ItemSlots     int `default:5`     // How many items can I carry?
	Vision        int `default:1`     // How FAR the entity can see, used in FOG removal
	Effects       string `default:"'` // Like attributes
}


/*func statCalc_Battle(off int, offMod int,offLevel int,def int, defMod, defLevel int) int{
	dmg := 1
	offense := off + (offMod * offLevel-1)
	defense := def + (defMod * defLevel-1)
	if(offense > defense){
		dmg = offense - def
	}
	return dmg
}*/
func statCalc_Battle(offStats *Stats, defStats *Stats) int{
	dmg := 1
	offense := offStats.GetOffenseWithMod()
	defense := defStats.GetDefenseWithMod()
	if(offense > defense){
		dmg = offense - defense
	}
	return dmg
}
//convert to %
func (s *Stats) UpdateHealth(delta int) {
	if s.checkEffects(overrides.ATTR_POISONOUS) && delta > 0 {
		delta = 0
	}
	s.Health += delta
	if (s.Health > s.GetHealthWithMod()) {
		s.Health = s.GetHealthWithMod()
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
	if (s.Mana > s.GetManaWithMod()) {
		s.Mana = s.GetManaWithMod()
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
/// GETTERS
func (s *Stats)IsPoisoned() bool {
	return s.checkEffects(overrides.ATTR_POISONOUS)
}
func (s *Stats)IsManaBurned() bool {
	return s.checkEffects(overrides.ATTR_MANABURN)
}
func (s *Stats)GetHealthWithMod() int {
	return s.MaxHealth + s.HealthMod*(s.Level-1) + s.ManaItemMod
}
func (s *Stats)GetManaWithMod() int {
	return s.MaxMana + s.ManaMod*(s.Level-1) + s.ManaItemMod
}
func (s *Stats)GetSpeedWithMod() int {
	return s.Speed + s.SpeedMod*(s.Level-1) + s.SpeedItemMod
}
func (s *Stats)GetOffenseWithMod() int {
	return s.Offense + s.OffMod*(s.Level-1) + s.OffItemMod
}
func (s *Stats)GetDefenseWithMod() int {
	return s.Defense + s.DefMod*(s.Level-1) + s.DefItemMod
}