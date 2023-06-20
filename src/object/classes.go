package object

type PlayerClass struct {
	Name string
	Stats Stats
	// TODO add starting items
}
var CLASS_PHYSICAL = PlayerClass{
	Name: "Laborer",
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 120,
		MaxMana:   110,
		Health:    120,
		Mana:      110,
		Defense:   1,
		Offense:   2,
		Speed:     1,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
var CLASS_MAGIC = PlayerClass{
	Name: "Inventor",
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 110,
		MaxMana:   120,
		Health:    110,
		Mana:      120,
		Defense:   2,
		Offense:   1,
		Speed:     1,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
var CLASS_SPEED = PlayerClass{
	Name: "Con Man",
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 100,
		MaxMana:   100,
		Health:    100,
		Mana:      100,
		Defense:   1,
		Offense:   1,
		Speed:     3,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    4,  // how FAR into fog we can see
	},
}