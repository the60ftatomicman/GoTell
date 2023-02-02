package tile

type Stats struct {
	MaxHealth int
	Health    int
	MaxMana   int
	Mana      int
	Level     int
	LevelMod  int
	XP        int
	Defense   int
	Offense   int
	Speed     int
	Favor     int
	FogRet    int
	Vision    int
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