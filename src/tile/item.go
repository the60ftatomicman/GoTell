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
		case "MANA":{p.Stats.UpdateMana(i.Delta)}
		case "OFFENSE":{p.Stats.Offense += i.Delta}	
	}
	return true;
}

func (i *Item) Convert(p *Player) {
	p.ChangeXP(i.ConversionPoints)
}

func generateItem(x int,y int, i Item) Item {
	i.X = x
	i.PrvX = x
	i.Y = y
	i.Prvy = y
	//i.Tile.Parent = &i //TODO - do we NEED the parent anymore? forgot it's purpose
	return i
}
//
//
//
//
func GenerateItemsFromFile() []Item{
	return []Item{
		generateItem(8 ,8,ITEM_HP),
		generateItem(9 ,8,ITEM_HP),
		generateItem(10,8,ITEM_MANA),
		generateItem(11,8,ITEM_HP),
		generateItem(12,8,ITEM_OFF_BOOST),
		generateItem(13,8,ITEM_HP),
	}
}
//
//
//
//
var ITEM_HP = Item{
	Name:      "HP Pot",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     25,
	ConversionPoints: 1,
	Affects:   "HEALTH",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      "H",
		Color:     core.TermCodes(core.FgWhite),
		BGColor:   core.TermCodes(core.BgRed),
		Attribute: core.ATTR_FOREGROUND+core.ATTR_ONETIME,
	},
}
var ITEM_MANA = Item{
	Name:      "Mana Pot",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     25,
	ConversionPoints: 1,
	Affects:   "MANA",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      "M",
		Color:     core.TermCodes(core.FgWhite),
		BGColor:   core.TermCodes(core.BgBlue),
		Attribute: core.ATTR_FOREGROUND+core.ATTR_ONETIME,
	},
}
var ITEM_OFF_BOOST = Item{
	Name:      "Pickaxe",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     5,
	ConversionPoints: 1,
	Affects:   "OFFENSE",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      "P",
		Color:     core.TermCodes(core.FgBlack),
		BGColor:   core.TermCodes(core.BgGreen),
		Attribute: core.ATTR_FOREGROUND,
	},
}