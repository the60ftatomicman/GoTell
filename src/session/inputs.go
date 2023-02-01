package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strconv"
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
			removeEnemy := s.Enemies[idx].Interaction(p)
			s.Profile.Health = strconv.Itoa(p.Stats.Health)
			s.Profile.Refresh() // would love to move to generic  session but I imagine this saves cycles
			s.Info.Message = "Enemy ["+s.Enemies[idx].Name+"] Health ["+strconv.Itoa(s.Enemies[idx].Stats.Health)+"]"
			if(removeEnemy){
				s.Info.Message = "Defeated ["+s.Enemies[idx].Name+"]"
				s.Screen.Buffer[p.Y][p.X].Pop()
				s.Enemies = append(s.Enemies[:idx], s.Enemies[idx+1:]...) 
			}
			s.Info.Refresh()
			s.Screen.Compile(&s.Profile,&s.Info)
		}
	}
	// -- movement otherwise
	nextTile := s.Screen.Buffer[p.Y][p.X].Get()
	if preventMovement(&p.Tile, &nextTile) {
		p.X = p.PrvX
		p.Y = p.PrvY
	}
	// -- now do the player placement
	s.Screen.Buffer[p.PrvY][p.PrvX].Pop()
	s.Screen.Buffer[p.Y][p.X].Set(p.Tile);
	// -- remove any fog
	fogRange := p.Stats.Vision
	for c := fogRange * -1; c <= fogRange; c++ {
		tileX := p.X+c;
		if(tileX < region.MAP_LEFT){tileX = region.MAP_LEFT}
		if(tileX > region.MAP_LEFT+region.MAP_COLUMNS){tileX = region.MAP_LEFT+region.MAP_COLUMNS}
		for r := fogRange * -1; r <= fogRange; r++ {
			tileY := p.Y+r;
			if(tileY < region.MAP_TOP){tileY = region.MAP_TOP}
			if(tileY > region.MAP_TOP+region.MAP_LINES){tileY = region.MAP_TOP+region.MAP_LINES}
			if (removeFog(&s.Screen.Buffer[tileY][tileX])) {
				p.UpdateHealth(p.Stats.FogRet)
				p.UpdateMana(p.Stats.FogRet)
				s.Profile.Health = strconv.Itoa(p.Stats.Health)
				s.Profile.Mana = strconv.Itoa(p.Stats.Mana)
				s.Profile.Refresh() // would love to move to generic  session but I imagine this saves cycles
				s.Screen.Compile(&s.Profile)
			}
		}
	}
	//some for loop here.
	
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

//TODO -- add all logic above into here. Also account for not being able to "see" through walls.
func removeFog(c *screen.Cell) bool {
	if (c.Get().Name == "FOG") {
		c.Pop()
		return true
	}
	return false
}