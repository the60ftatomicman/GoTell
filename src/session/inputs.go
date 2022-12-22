package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strings"
)

//Input_Buffer = [4]string{"","","",""}
//var Input_nextIdx int = 0

func handleInput(input string, p *tile.Player, s *screen.Screen) bool {
	p.PrvY = p.Y
	p.PrvX = p.X
	switch input {
	case "Q":
		{
			return true
		}
	case "w":
		{
			if p.Y > 1 {
				p.Y -= 1
			}
		}
	case "d":
		{
			if p.X < region.MAP_COLUMNS-1 {
				p.X += 1
			}
		}
	case "s":
		{
			if p.Y < region.MAP_LINES-1 {
				p.Y += 1
			}
		}
	case "a":
		{
			if p.X > 1 {
				p.X -= 1
			}
		}
	}
	// Test if we are fighting
	objTile := s.Buffer[p.Y][p.X]
	objEnemy := performCombat(&p.Tile, &objTile)
	if objEnemy != nil && objEnemy.Interaction() {
		s.Set(objEnemy.Tile, objEnemy.X, objEnemy.Y)
	}
	//
	if preventMovement(&p.Tile, &objTile) {
		p.X = p.PrvX
		p.Y = p.PrvY
		s.Set(p.Tile, p.X, p.Y)
	} else {
		s.Set(p.UnderTile, p.PrvX, p.PrvY)
		p.UnderTile = s.Buffer[p.Y][p.X]
		s.Set(p.Tile, p.X, p.Y)
	}

	return false
}

func performCombat(tA *tile.Tile, tB *tile.Tile) *tile.Enemy {
	if strings.Contains(tB.Attribute, core.ATTR_FIGHTABLE) {
		return tB.Parent
	}
	return nil
}

func preventMovement(tA *tile.Tile, tB *tile.Tile) bool {
	var prevent bool = false
	if strings.Contains(tA.Attribute, core.ATTR_SOLID) && strings.Contains(tB.Attribute, core.ATTR_SOLID) {
		prevent = true
	}

	return prevent
}
