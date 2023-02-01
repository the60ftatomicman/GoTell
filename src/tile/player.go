package tile

import "example/gotell/src/core"

type Player struct {
	UnderTile        Tile
	Tile             Tile
	Name             string
	X, Y, PrvX, PrvY int
	Stats            Stats
}

func GeneratePlayer() Player {
	return Player{
		X:    1,
		Y:    5,
		PrvX: 1,
		PrvY: 5,
		Stats: Stats{
			Health: 100,
			Defense: 1,
			Offense: 1,
			Speed:   2,
			FogRet:  10, // how much MANA and HEALTH we get back when uncovering FOG
			Vision:  3,  // how FAR into fog we can see
		},
		Tile: Tile{
			Name:      "PLAYER",
			Icon:      core.Icons(core.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: core.ATTR_SOLID + core.ATTR_FOREGROUND,
		},
	}
}

func (p *Player) UpdateHealth(delta int) {
	p.Stats.Health += delta
	if (p.Stats.Health > 100) {
		p.Stats.Health = 100
	}
	if (p.Stats.Health < 0) {
		p.Stats.Health = 0
	}
}

func (p *Player) UpdateMana(delta int) {
	p.Stats.Mana += delta
	if (p.Stats.Mana > 100) {
		p.Stats.Mana = 100
	}
	if (p.Stats.Mana < 0) {
		p.Stats.Mana = 0
	}
}