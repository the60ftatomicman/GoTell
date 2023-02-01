package tile

type Stats struct {
	Health  int
	Mana    int
	Defense int
	Offense int
	Speed   int
	Favor   int
	XP      int
	FogRet  int
	Vision  int
}

func statCalc_Battle(off int, def int) int{
	dmg := 0
	if(off > def){
		dmg = off - def
	}
	return dmg
}