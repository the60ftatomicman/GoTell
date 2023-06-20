package object

import (
	"example/gotell/src/core"
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
)

type Player struct {
	Tile             tile.Tile
	Name             string `default:""`
	Class            string `default:""`
	Dir              string `default:""`
	X,Y,DirX,DirY    int    `default:0`
	Stats            Stats
	Items            []Item
}

func GeneratePlayer() Player {
	return Player{
		Name: "Billsy",
		Class: "Unmolded",
		X:    1,
		Y:    5,
		Stats: Stats{
			Level:     1,
			LevelMod:  10,
			MaxHealth: 100,
			MaxMana:   100,
			Health:    100,
			Mana:      100,
			Defense:  1,
			Offense:  1,
			Speed:    2,
			FogRet:   10, // how much MANA and HEALTH we get back when uncovering FOG
			Vision:   3,  // how FAR into fog we can see
		},
		Tile: tile.Tile{
			Name:      "PLAYER",
			Icon:      tile.Icons(overrides.ICON_PLAYER),
			Color:     core.TermCodes(core.FgGreen),
			Attribute: tile.GenerateAttributes(overrides.ATTR_SOLID),
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

func (p *Player) GetDirString() string{
	dir := ""
	if (p.DirY < 0){dir+="NORTH"}
	if (p.DirY > 0){dir+="SOUTH"}
	if (p.DirX < 0){dir+="WEST"}
	if (p.DirX > 0){dir+="EAST"}
	return dir
}