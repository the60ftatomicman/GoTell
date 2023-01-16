package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strings"
)

//Input_Buffer = [4]string{"","","",""}
//var Input_nextIdx int = 0

func hanleInputStateSwitching(input string, s *Session, inf *region.Info) bool{
	switch input {
		case "Q":
			{
				return true
			}
		case "i":
			{
				if s.State != STATE_INVENTORY{
					s.State = STATE_INVENTORY
					inf.Message = "Currently [Invetory]: 1-3 (item), switch to (m)oving"
					inf.Refresh()
				}
			}
		case "m":
			{
				if s.State != STATE_MOVING{
					s.State = STATE_MOVING
					inf.Message = "Currently [MOVING]: WASD (moves), switch to (i)nventory, (Q)uit"
					inf.Refresh()
				}
			}
		case "1":
			{
				if s.State == STATE_INVENTORY{
					s.State = STATE_ITEM
					inf.Message = "Looking at [some item]: (D)rop,(U)se, switch to (i)nventory, switch to (m)oving"
					inf.Refresh()
				}
			}
	}
	return false;
}

func handleInputMoving(input string, p *tile.Player, s *Session) {
	p.PrvY = p.Y
	p.PrvX = p.X
	switch input {
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
	for idx,enemy := range s.Enemies {
		if enemy.X == p.X && enemy.Y == p.Y {
			enemyObj := s.Enemies[idx]
			removeEnemy := enemyObj.Interaction()
			if(removeEnemy){
				 s.Screen.Buffer[p.Y][p.X].Pop()
				 s.Enemies = append(s.Enemies[:idx], s.Enemies[idx+1:]...)
			}
		}
	}
	// -- movement otherwise
	nextTile := s.Screen.Buffer[p.Y][p.X].Get()
	if preventMovement(&p.Tile, &nextTile) {
		p.X = p.PrvX
		p.Y = p.PrvY
	}
	s.Screen.Buffer[p.PrvY][p.PrvX].Pop()
	s.Screen.Buffer[p.Y][p.X].Set(p.Tile);
	
}

func handleInputInventory(input string, p *region.Profile,inf *region.Info){
	switch input {
		case "1":
			{

			}
		case "2":
			{

			}
	}
}

func preventMovement(tA *tile.Tile, tB *tile.Tile) bool {
	var prevent bool = false
	if strings.Contains(tA.Attribute, core.ATTR_SOLID) && strings.Contains(tB.Attribute, core.ATTR_SOLID) {
		prevent = true
	}

	return prevent
}
