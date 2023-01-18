package tile

type Stats struct {
	Health  int
	Defense int
	Offense int
	Speed   int
}

func statCalc_Battle(off int, def int) int{
	dmg := 0
	if(off > def){
		dmg = off - def
	}
	return dmg
}