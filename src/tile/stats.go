package tile

type Stats struct {
	MaxHealth int `default:100`
	Health    int `default:100`
	MaxMana   int `default:100`
	Mana      int `default:100`
	Level     int `default:1`
	LevelMod  int `default:10`
	XP        int `default:0`
	Defense   int `default:1`
	Offense   int `default:1`
	Speed     int `default:1`
	Favor     int `default:0`
	FogRet    int `default:25`
	Vision    int `default:1`
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