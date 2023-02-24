package session

import (
	"example/gotell/src/screen/region"
)

func handleGlobalStateSwitching(input string, s *Session) bool{
	switch input {
		case "Q":
			{
				return true
			}
	}
	return false
}
func handleInputStateSwitching(input string, s *Session) bool{
	previousState := s.State.Name
	switch input {
		case "i":
			{
				if s.State.Name != STATE_INVENTORY.Name{
					s.State = STATE_INVENTORY
				}
			}
		case "m":
			{
				if s.State.Name != STATE_MOVING.Name{
					s.State = STATE_MOVING
					s.Profile.SelectedItem = ""
				}
			}
		case "p":{
			if s.State.Name == STATE_MOVING.Name {
				s.State = STATE_GETITEM
			}
		}
		case "r":{
			//	DEBUG ONLY TO REVIVE MYSELF!
			if s.State.Name == STATE_MOVING.Name {
				s.State = STATE_POPUP
			}
		}
		default: {
			//For passthrough states like picking up an item
			switch s.State.Name {
				case STATE_GETITEM.Name: {s.State = STATE_MOVING}
			}
		}
	}
	return previousState != s.State.Name
}

///
/// TODO -- this needs to go back into movement I think
///

//TODO -- use this more!
func getTileXY(playerX int,playerY int,colDelta int,rowDelta int) (int,int) {
		tileX := playerX+colDelta;
		if(tileX < region.MAP_LEFT){tileX = region.MAP_LEFT}
		if(tileX > region.MAP_LEFT+region.MAP_COLUMNS){tileX = region.MAP_LEFT+region.MAP_COLUMNS}
		tileY := playerY+rowDelta;
		if(tileY < region.MAP_TOP){tileY = region.MAP_TOP}
		if(tileY > region.MAP_TOP+region.MAP_LINES){tileY = region.MAP_TOP+region.MAP_LINES}
		return tileX,tileY
}

//TODO -- Account for not being able to "see" through walls.
func removeFog(s *Session,colDelta int,rowDelta int) string{
	p := &s.Player
	tileX,tileY := getTileXY(p.X,p.Y,colDelta,rowDelta)
	s.Screen.Buffer[tileY][tileX].Get()
	//TODO -- hate this is hard coded
	if (s.Screen.Buffer[tileY][tileX].Get().Name  == "FOG") {
		s.Screen.Buffer[tileY][tileX].Pop()
		p.Stats.UpdateHealth(p.Stats.FogRet)
		p.Stats.UpdateMana(p.Stats.FogRet)
		// Update all those enemies health (dun dun dun)
		for idx := range s.Level.Enemies {
			e := &s.Level.Enemies[idx]
			e.Stats.UpdateHealth(e.Stats.FogRet)
			e.Stats.UpdateMana(e.Stats.FogRet)
		}
	}
	//return the value of the current tile
	return s.Screen.Buffer[tileY][tileX].Get().Attribute
}