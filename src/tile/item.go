package tile

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
	if i.Uses > 0  || i.Uses == UNLIMITED_USES {
		switch(i.Affects){
			case Affects(Health) :{s.UpdateHealth(i.Delta)}
			case Affects(Mana)   :{s.UpdateMana(i.Delta)  }
			case Affects(Offense):{s.Offense += i.Delta   }
			case Affects(Defense):{s.Defense += i.Delta   }
			case Affects(Speed)  :{s.Speed += i.Delta     }	
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
	Delta:     5,
	Uses:      1,
	ConversionPoints: 10,
	Affects:   Affects(Offense),
	Tile: EQUIPTMENT,
}
var ITEM_DEF_BOOST = Item{
	Name:      "Hard Hat",
	X:         0,
	Y:         0,
	Delta:     5,
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