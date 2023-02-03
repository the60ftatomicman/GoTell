package tile

import "example/gotell/src/core"

type Player struct {
	UnderTile        Tile
	Tile             Tile
	Name             string
	X, Y, PrvX, PrvY,DirX,DirY int
	Stats            Stats
}

func GeneratePlayer() Player {
	return Player{
		X:    1,
		Y:    5,
		PrvX: 1,
		PrvY: 5,
		DirX: 0,
		DirY: 0,
		Stats: Stats{
			Level:     1,
			MaxHealth: 100,
			Health:   100,
			Defense:  1,
			Offense:  1,
			Speed:    2,
			FogRet:   10, // how much MANA and HEALTH we get back when uncovering FOG
			Vision:   3,  // how FAR into fog we can see
		},
		Tile: Tile{
			Name:      "PLAYER",
			Icon:      core.Icons(core.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: core.ATTR_SOLID + core.ATTR_FOREGROUND,
		},
	}
}

func (p *Player)GetViewRanges() (int,int,int,int,int,int){
	fogRange := p.Stats.Vision
	xStart := fogRange * -1
	xEnd   := fogRange
	xInc   := 1
	if(p.DirX != 0){
		xStart = 0
		xEnd = fogRange * p.DirX
		xInc = p.DirX
	}
	yStart := fogRange * -1
	yEnd   := fogRange
	yInc   := 1
	if(p.DirY != 0){
		yStart = 0
		yEnd = fogRange * p.DirY
		yInc = p.DirY
	}
	return xStart,xEnd,xInc,yStart,yEnd,yInc
}