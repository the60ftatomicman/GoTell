package object

import (
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strconv"
	"strings"
)

type Affects string
const (
	Health     Affects = "HEALTH"
	MaxHealth          = "HEALTHCAP"
	Mana               = "MANA"
	MaxMana            = "MANACAP"
	Offense            = "OFFENSE"
	Defense            = "DEFENSE"
	Speed              = "SPEED"
)
const UNLIMITED_USES = -1 //TODO -- do we use this anymore?

type Item struct {
	Tile             tile.Tile           `default:"Unknown Item"`
	Name             string              `default:"OK"`
	Description      string              `default:""`
	X,Y,Cost,Delta,ConversionPoints int  `default:0`
	Uses             int                 `default:1`
	Affects          Affects
}

func (i *Item) Interaction(s *Stats) bool{
	//TODO -- is the ues code actually used or did I fix this with attributes?
	if i.Uses > 0  || i.Uses == UNLIMITED_USES {
		switch(i.Affects){
			case Affects(Health) :{
				s.RemoveEffects(overrides.ATTR_POISONOUS)
				s.UpdateHealth(i.Delta)
			}
			case Affects(Mana)   :{
				s.RemoveEffects(overrides.ATTR_MANABURN)
				s.UpdateMana(i.Delta)  
			}
			case Affects(MaxHealth):{s.HealthItemMod += i.Delta }
			case Affects(MaxMana):{s.ManaItemMod     += i.Delta }
			case Affects(Offense):{s.OffItemMod      += i.Delta }
			case Affects(Defense):{s.DefItemMod      += i.Delta }
			case Affects(Speed)  :{s.SpeedItemMod    += i.Delta }	
		}
		if i.Uses != UNLIMITED_USES{
			i.Uses -= 1
		}
		return i.Uses == 0
	}else{
		return true;
	}
}

func (i *Item) Convert(s *Stats) {
	s.ChangeXP(i.ConversionPoints)
}
func (i *Item) GetBufferData() (int,int,string,tile.Tile) {
	return i.Y,i.X,i.Name,i.Tile
}

func generateItem(x int,y int, i Item) Item {
	i.X = x
	i.Y = y
	return i
}
//
//
//
//
func GenerateItemsFromFile(fileData []string) []Item{
	itemList := []Item{}
	for _,row := range fileData {
		itemList = append(itemList,fileParserItem(row)...)
	}
	return itemList
}

var dataConverterItem = map[string]Item{
     "ITEM_HP": ITEM_HP,
     "ITEM_MANA": ITEM_MANA,
	 "ITEM_OFF_BOOST":ITEM_OFF_BOOST,
	 "ITEM_DEF_BOOST":ITEM_DEF_BOOST,
	 "ITEM_SPELL_DMG":ITEM_SPELL_DMG,
}

func fileParserItem(itemVals string) ([]Item){
	items    := []Item{}
	countVal := strings.Split(itemVals, ":")
	count,_  := strconv.Atoi(countVal[0])
	value    := countVal[1]
	for len(items) < count {
		items  = append(items, dataConverterItem[value])
	}
	return items
}
//
//
//
//
//
var ITEM_HP = Item{
	Name:      "HP Pot",
	Description: "+25 HP",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     25,
	ConversionPoints: 3,
	Affects:   Affects(Health),
	Tile: overrides.POTION_HEALTH,
}
var ITEM_MANA = Item{
	Name:      "Mana Pot",
	Description: "+25 Mana",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     3,
	ConversionPoints: 1,
	Affects:  Affects(Mana),
	Tile: overrides.POTION_MANA,
}
var ITEM_OFF_BOOST = Item{
	Name:      "Pickaxe",
	Description: "+5 Offense",
	X:         0,
	Y:         0,
	Delta:     5,
	Uses:      UNLIMITED_USES,
	ConversionPoints: 5,
	Affects:   Affects(Offense),
	Tile: overrides.EQUIPTMENT,
}
var ITEM_DEF_BOOST = Item{
	Name:      "Hard Hat",
	Description: "+5 Defense",
	X:         0,
	Y:         0,
	Delta:     5,
	Uses:      UNLIMITED_USES,
	ConversionPoints: 5,
	Affects:   Affects(Defense),
	Tile: overrides.EQUIPTMENT,
}
var ITEM_SPEED_BOOST = Item{
	Name:      "Roller Blades",
	Description: "+5 Speed",
	X:         0,
	Y:         0,
	Uses:      UNLIMITED_USES,
	Delta:     5,
	ConversionPoints: 5,
	Affects:   Affects(Speed),
	Tile: overrides.EQUIPTMENT,
}
var ITEM_SPELL_DMG = Item{
	Name:      "Moose shot",
	Description: "gives 30 damage",
	X:         0,
	Y:         0,
	Uses:      UNLIMITED_USES,
	Cost:      -30,
	Delta:     -30,
	ConversionPoints: 10,
	Affects:   Affects(Health),
	Tile: overrides.SPELL,
}