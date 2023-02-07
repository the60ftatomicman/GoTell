package tile


type Item struct {
	Tile             Tile                `default:"Unknown Item"`
	Name             string              `default:"OK"`
	X,Y,Cost,Delta,ConversionPoints int  `default:0`
	Uses             int                 `default:1`
	Affects          string              `default:""` //TODO -- make ENUM
}

var UNLIMITED_USES = -1

//DUN DUN DUN, we are 100 going to 
func (i *Item) Interaction(s *Stats) bool{
	if i.Uses > 0  || i.Uses == UNLIMITED_USES {
		switch(i.Affects){
			case "HEALTH":{s.UpdateHealth(i.Delta)}
			case "MANA":{s.UpdateMana(i.Delta)}
			case "OFFENSE":{s.Offense += i.Delta}
			case "DEFENSE":{s.Defense += i.Delta}
			case "SPEED":{s.Speed += i.Delta}	
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
func GenerateItemsFromFile() []Item{
	return []Item{
		generateItem(8 ,8,ITEM_HP),
		generateItem(9 ,8,ITEM_HP),
		generateItem(10,8,ITEM_MANA),
		generateItem(11,8,ITEM_DEF_BOOST),
		generateItem(12,8,ITEM_SPELL_DMG),
		//generateItem(12,8,ITEM_OFF_BOOST),
		//generateItem(13,8,ITEM_SPEED_BOOST),
	}
}
//
//
// TODO add "item" tiles to tiles file
//
var ITEM_HP = Item{
	Name:      "HP Pot",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     25,
	ConversionPoints: 1,
	Affects:   "HEALTH",
	Tile: POTION_HEALTH,
}
var ITEM_MANA = Item{
	Name:      "Mana Pot",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     25,
	ConversionPoints: 1,
	Affects:   "MANA",
	Tile: POTION_MANA,
}
var ITEM_OFF_BOOST = Item{
	Name:      "Pickaxe",
	X:         0,
	Y:         0,
	Delta:     5,
	Uses:      1,
	ConversionPoints: 10,
	Affects:   "OFFENSE",
	Tile: EQUIPTMENT,
}
var ITEM_DEF_BOOST = Item{
	Name:      "Hard Hat",
	X:         0,
	Y:         0,
	Delta:     5,
	Uses:      1,
	ConversionPoints: 10,
	Affects:   "DEFENSE",
	Tile: EQUIPTMENT,
}
var ITEM_SPEED_BOOST = Item{
	Name:      "Roller Blades",
	X:         0,
	Y:         0,
	Uses:      1,
	Delta:     5,
	ConversionPoints: 10,
	Affects:   "SPEED",
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
	Affects:   "HEALTH",
	Tile: SPELL,
}