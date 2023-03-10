package tile

import (
	"example/gotell/src/core"
	"strconv"
	"strings"
)

type Affects string
const (
	Health     Affects = "HEALTH"
	Mana               = "MANA"
	Offense            = "OFFENSE"
	Defense            = "DEFENSE"
	Speed              = "SPEED"
)
const UNLIMITED_USES = -1 //TODO -- do we use this anymore?

type Item struct {
	Tile             Tile                `default:"Unknown Item"`
	Name             string              `default:"OK"`
	X,Y,Cost,Delta,ConversionPoints int  `default:0`
	Uses             int                 `default:1`
	Affects          Affects
}

func (i *Item) Interaction(s *Stats) bool{
	//TODO -- is the ues code actually used or did I fix this with attributes?
	if i.Uses > 0  || i.Uses == UNLIMITED_USES {
		switch(i.Affects){
			case Affects(Health) :{
				s.RemoveEffects(core.ATTR_POISONOUS)
				s.UpdateHealth(i.Delta)
			}
			case Affects(Mana)   :{
				s.RemoveEffects(core.ATTR_MANABURN)
				s.UpdateMana(i.Delta)  
			}
			case Affects(Offense):{s.Offense += i.Delta }
			case Affects(Defense):{s.Defense += i.Delta }
			case Affects(Speed)  :{s.Speed   += i.Delta }	
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
func (i *Item) GetBufferData() (int,int,string,Tile) {
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
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     25,
	ConversionPoints: 1,
	Affects:   Affects(Health),
	Tile: POTION_HEALTH,
}
var ITEM_MANA = Item{
	Name:      "Mana Pot",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     25,
	ConversionPoints: 1,
	Affects:  Affects(Mana),
	Tile: POTION_MANA,
}
var ITEM_OFF_BOOST = Item{
	Name:      "Pickaxe",
	X:         0,
	Y:         0,
	Delta:     1,
	Uses:      1,
	ConversionPoints: 10,
	Affects:   Affects(Offense),
	Tile: EQUIPTMENT,
}
var ITEM_DEF_BOOST = Item{
	Name:      "Hard Hat",
	X:         0,
	Y:         0,
	Delta:     1,
	Uses:      1,
	ConversionPoints: 10,
	Affects:   Affects(Defense),
	Tile: EQUIPTMENT,
}
var ITEM_SPEED_BOOST = Item{
	Name:      "Roller Blades",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     5,
	ConversionPoints: 10,
	Affects:   Affects(Speed),
	Tile: EQUIPTMENT,
}
var ITEM_SPELL_DMG = Item{
	Name:      "Moose shot",
	X:         0,
	Y:         0,
	Uses:      UNLIMITED_USES,
	Cost:      -100,
	Delta:     -5,
	ConversionPoints: 10,
	Affects:   Affects(Health),
	Tile: SPELL,
}