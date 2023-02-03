package tile

import "example/gotell/src/core"


type Item struct {
	Tile             Tile
	Name             string
	X, Y, PrvX, Prvy,Delta,ConversionPoints int
	Affects          string //TODO -- make ENUM
}

func (i *Item) Interaction(p *Player) bool{
	switch(i.Affects){
		case "HEALTH":{p.Stats.UpdateHealth(i.Delta)}	
	}
	return true;
}

func (i *Item) Convert(p *Player) {
	p.ChangeXP(i.ConversionPoints)
}

func generateItem(x int,y int) Item {
	i := Item{
		Name:      "HP Pot",
		X:         x,
		Y:         y,
		PrvX:      x,
		Prvy:      y,
		Delta:     25,
		Affects:   "HEALTH",
	}
	i.Tile = Tile{
		Name:      "ITEM",
		Icon:      "H",
		Color:     core.TermCodes(core.FgWhite),
		BGColor:   core.TermCodes(core.BgRed),
		Attribute: core.ATTR_FOREGROUND,
		Parent:    &i,
	}
	return i
}
//
//
//
//
func GenerateItemsFromFile() []Item{
	return []Item{
		generateItem(8,8),
		generateItem(9,8),
		generateItem(10,8),
		generateItem(11,8),
		generateItem(12,8),
		generateItem(13,8),
	}
}