package tile

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
}

func statCalc_Battle(off int, def int, offMod int) int{
	dmg := 0
	if((off * offMod) > def){
		dmg = (off * offMod) - def
	}
	return dmg
}

func (s *Stats) UpdateHealth(delta int) {
	s.Health += delta
	if (s.Health > s.MaxHealth) {
		s.Health = s.MaxHealth
	}
	if (s.Health < 0) {
		s.Health = 0
	}
}

func (s *Stats) UpdateMana(delta int) {
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
	if(s.XP >= 10) {
		s.Level  = (s.XP / 10) + 1
		s.XP     = s.XP % 10
		s.Health = s.MaxHealth
		s.Mana   = s.MaxMana
	}
}