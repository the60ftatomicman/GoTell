package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strconv"
	"strings"
)

func hanleInputStateSwitching(input string, s *Session) bool{
	switch input {
		case "Q":
			{
				return true
			}
		case "i":
			{
				if s.State != STATE_INVENTORY{
					s.State = STATE_INVENTORY
				}
			}
		case "m":
			{
				if s.State != STATE_MOVING{
					s.State = STATE_MOVING
				}
			}
		case "r":
			{
				if s.State == STATE_DEAD{
					s.Player.Stats.UpdateHealth(s.Player.Stats.MaxHealth)
					s.Player.Stats.UpdateMana(s.Player.Stats.MaxMana)
					s.State = STATE_MOVING
					s.Info.Set("Currently [MOVING]: WASD (moves), switch to (i)nventory, (Q)uit")
					s.Info.Refresh()
					s.Profile.Health = strconv.Itoa(s.Player.Stats.Health)
					s.Profile.Mana = strconv.Itoa(s.Player.Stats.Mana)
					s.Profile.Refresh()
				}
			}
		case "1":
			{
				if s.State == STATE_INVENTORY{
					s.State = STATE_ITEM
				}
			}
	}
	return false;
}

///
///
///
func preventMovement(tA *tile.Tile, tB *tile.Tile) bool {
	var prevent bool = false
	if strings.Contains(tA.Attribute, core.ATTR_SOLID) && strings.Contains(tB.Attribute, core.ATTR_SOLID) {
		prevent = true
	}

	return prevent
}

func getTileXY(playerX int,playerY int,colDelta int,rowDelta int) (int,int) {
		tileX := playerX+colDelta;
		if(tileX < region.MAP_LEFT){tileX = region.MAP_LEFT}
		if(tileX > region.MAP_LEFT+region.MAP_COLUMNS){tileX = region.MAP_LEFT+region.MAP_COLUMNS}
		tileY := playerY+rowDelta;
		if(tileY < region.MAP_TOP){tileY = region.MAP_TOP}
		if(tileY > region.MAP_TOP+region.MAP_LINES){tileY = region.MAP_TOP+region.MAP_LINES}
		return tileX,tileY
}

//TODO -- add all logic above into here. Also account for not being able to "see" through walls.
//TODo -- why am I passing the player....
func removeFog(s *Session,colDelta int,rowDelta int) string{
	p := &s.Player
	tileX,tileY := getTileXY(p.X,p.Y,colDelta,rowDelta)
	s.Screen.Buffer[tileY][tileX].Get()
	if (s.Screen.Buffer[tileY][tileX].Get().Name  == "FOG") {
		s.Screen.Buffer[tileY][tileX].Pop()
		p.Stats.UpdateHealth(p.Stats.FogRet)
		p.Stats.UpdateMana(p.Stats.FogRet)
		s.Profile.Health = strconv.Itoa(p.Stats.Health)
		s.Profile.Mana = strconv.Itoa(p.Stats.Mana)
		// Update all those enemies health!
		for idx,_ := range s.Enemies {
			e := &s.Enemies[idx]
			e.Stats.UpdateHealth(e.Stats.FogRet)
			e.Stats.UpdateMana(e.Stats.FogRet)
		}
	}
	//return the value of the current tile
	return s.Screen.Buffer[tileY][tileX].Get().Attribute
}