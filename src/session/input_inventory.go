package session

import (
	"example/gotell/src/core/tile"
	overrides "example/gotell/src/core_overrides"
	"strconv"
)


func handleInputInventory(input string, s *Session) bool{
	s.Info.Set(MENU_INVENTORY(&s.Profile,"")...)
	idx,notInt := strconv.Atoi(input)
	if(notInt == nil){
		if(idx > 0 && idx <= len(s.Profile.Player.Items)){
			s.Profile.SelectedItem = input
			if(tile.CheckAttributes(s.Profile.Player.Items[idx-1].Tile,overrides.ATTR_SPELL)){
				s.State = STATE_SPELL
			}else{
				s.State = STATE_ITEM
			}
			return true
		}
	}else{
		s.Info.Set(MENU_INVENTORY(&s.Profile,input)...)	
	}
	return false
}