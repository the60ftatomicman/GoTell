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
		case "DEFENSE":{p.Stats.Defense += i.Delta}
		case "SPEED":{p.Stats.Speed += i.Delta}	
	}
	return true;
}

func (i *Item) Convert(p *Player) {
	p.ChangeXP(i.ConversionPoints)
}
func (i *Item) GetBufferData() (int,int,string,Tile) {
	return i.Y,i.X,i.Name,i.Tile
}

func generateItem(x int,y int, i Item) Item {
	i.X = x
	i.PrvX = x
	i.Y = y
	i.Prvy = y
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
		generateItem(11,8,ITEM_DEF_BOOST),
		generateItem(12,8,ITEM_OFF_BOOST),
		generateItem(13,8,ITEM_SPEED_BOOST),
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
		Icon:      core.ICON_HEALTH,
		Color:     core.TermCodes(core.FgWhite),
		BGColor:   core.TermCodes(core.BgRed),
		Attribute: core.ATTR_ONETIME,
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
		Icon:      core.ICON_MANA,
		Color:     core.TermCodes(core.FgWhite),
		BGColor:   core.TermCodes(core.BgBlue),
		Attribute: core.ATTR_ONETIME,
	},
}
var ITEM_OFF_BOOST = Item{
	Name:      "Pickaxe",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     5,
	ConversionPoints: 10,
	Affects:   "OFFENSE",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      core.ICON_OFF_BOOST,
		Color:     core.TermCodes(core.FgBlack),
		BGColor:   core.TermCodes(core.BgGreen),
		Attribute: "",
	},
}
var ITEM_DEF_BOOST = Item{
	Name:      "Hard Hat",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     5,
	ConversionPoints: 10,
	Affects:   "DEFENSE",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      core.ICON_OFF_BOOST,
		Color:     core.TermCodes(core.FgBlack),
		BGColor:   core.TermCodes(core.BgGreen),
		Attribute: "",
	},
}
var ITEM_SPEED_BOOST = Item{
	Name:      "Roller Blades",
	X:         0,
	Y:         0,
	PrvX:      0,
	Prvy:      0,
	Delta:     5,
	ConversionPoints: 10,
	Affects:   "SPEED",
	Tile: Tile{
		Name:      "ITEM",
		Icon:      core.ICON_OFF_BOOST,
		Color:     core.TermCodes(core.FgBlack),
		BGColor:   core.TermCodes(core.BgGreen),
		Attribute: "",
	},
}