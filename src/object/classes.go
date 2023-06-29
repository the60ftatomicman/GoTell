package object

type PlayerClass struct {
	Name string
	Description []string
	Stats Stats
	// TODO add starting items
}
var CLASS_PHYSICAL = PlayerClass{
	Name: "Laborer",
	Description: []string{
		"Strong! Dependable! Rugged!",
		"Labrorer is the choice for",
		"those who WANT to just use brute",
		"strength to just clobber their way through the mine.",
	},
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 100,
		HealthMod: 10,
		Health:    100,
		MaxMana:   30,
		ManaMod:   5,
		Mana:      30,
		Defense:   10,
		DefMod:    10,
		Offense:   10,
		OffMod:    10,
		Speed:     2,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
var CLASS_MAGIC = PlayerClass{
	Name: "Inventor",
	Description: []string{
		"Seeking fame; Inventors find themselves in the mine to prove",
		"that their engineering prowress is unmatched. What does",
		"that egghead Edison or timid Telsa have that they don't?",
		"Pick if your into indirect combat.",
	},
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 100,
		HealthMod: 10,
		Health:    100,
		MaxMana:   60,
		ManaMod:   5,
		Mana:      60,
		Defense:   9,
		DefMod:    9,
		Offense:   7,
		OffMod:    7,
		Speed:     1,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
var CLASS_SPEED = PlayerClass{
	Name: "Con Man",
	Description: []string{
		"\"I don't know how I got here but I heard there's gold.\"",
		"The con man isn't strong, or smart. They are fast however",
		"and able to detect things a bit further away than most.",
	},
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 80,
		HealthMod: 10,
		Health:    80,
		MaxMana:   45,
		ManaMod:   5,
		Mana:      45,
		Defense:   9,
		DefMod:    9,
		Offense:   9,
		OffMod:    9,
		Speed:     3,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
//UNTESTED!
var CLASS_EXP = PlayerClass{
	Name: "Scholar",
	Description: []string{
		"Reports of monsters and new geological formations have piqued",
		"the interest of scholars from all over.",
		"They start weak but learn twice as fast as the other classes.",
	},
	Stats: Stats{
		Level:     1,
		LevelMod:  5,
		MaxHealth: 50,
		HealthMod: 10,
		Health:    50,
		MaxMana:   30,
		ManaMod:   5,
		Mana:      30,
		Defense:   6,
		DefMod:    11,
		Offense:   6,
		OffMod:    11,
		Speed:     1,
		FogRet:    25, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    3,  // how FAR into fog we can see
	},
}
var CLASS_FOG = PlayerClass{
	Name: "Detective",
	Description: []string{
		"The GOVERNMENT thinks they need to snope on our business here.",
		"Detectives and agents are being sent from the 3 letters to check", // THIS IS OUR LONGEST! nothing more nothing less....
		"out what we're doing. These goons gain their stats back twice",
		"as fast as the others.Uncovering is key",
	},
	Stats: Stats{
		Level:     1,
		LevelMod:  10,
		MaxHealth: 90,
		HealthMod: 9,
		Health:    90,
		MaxMana:   30,
		ManaMod:   5,
		Mana:      30,
		Defense:   9,
		DefMod:    9,
		Offense:   9,
		OffMod:    9,
		Speed:     2,
		FogRet:    50, // how much MANA and HEALTH we get back when uncovering FOG
		Vision:    4,  // how FAR into fog we can see
	},
}