package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strconv"
)

//TODO -- make a MENU for this.
func handleGetItem(input string, s *Session) bool{
	msg := "you are NOT on an item."
	if(len(s.Player.Items) < region.LINE_VAR_ITEM_COUNT){
		for idx,item := range s.Level.Items {
			if (item.X == s.Player.X && item.Y == s.Player.Y){
				s.Player.Items = append(s.Player.Items,item)
				msg = "Picked up ["+item.Name+"]"
				//Remove player and item
				s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
				s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
				s.Screen.Buffer[s.Player.Y][s.Player.X].Set(s.Player.Tile)
				s.Level.Items = append(s.Level.Items[:idx], s.Level.Items[idx+1:]...)
				if(tile.CheckAttributes(item.Tile,core.ATTR_EQUIPTABLE)){
					item.Interaction(&s.Player.Stats)
				}
			}
		}
	}else{
		msg = "Your inventory is FULL"
	}
	s.Info.Set(msg,"Press [wasd] to move")
	s.State = STATE_MOVING
	return false
}
func handleItemAction(input string, s *Session) bool{
	//s.State = STATE_INVENTORY
	return false
}
func handleInputItem(input string, s *Session) bool{
	idx,notInt    := strconv.Atoi(s.Profile.SelectedItem)
	secondRefresh := false
	if notInt == nil {
		//TODO -- error handling
		idx            = idx - 1
		item          := s.Player.Items[idx]
		s.Info.Set(MENU_ITEM(&item)...)
		switch input {
			case "u":
				{
					if(tile.CheckAttributes(item.Tile,core.ATTR_SPELL)){
						s.State = STATE_SPELL
						//TODO -- this is a hack for now.
						handleInputSpell(input, s)
					}else{
						if(item.Interaction(&s.Player.Stats)){
							s.Player.Items = append(s.Player.Items[:idx], s.Player.Items[idx+1:]...)
							s.Profile.SelectedItem = ""
							secondRefresh = true
							s.State = STATE_MOVING
						}
					}
				}
			case "c":
				{
					s.Player.Stats.ChangeXP(item.ConversionPoints)
					s.Player.Items = append(s.Player.Items[:idx], s.Player.Items[idx+1:]...)
					s.Profile.SelectedItem = ""
					secondRefresh = true
					s.State = STATE_ITEMACTION
				}
			case "d":
				{
					if(!tile.CheckAttributes(item.Tile,core.ATTR_EQUIPTABLE)){
						item.Delta *= -1
						item.Interaction(&s.Player.Stats)
					}
					item.Delta *= -1
					item.X = s.Player.X
					item.Y = s.Player.Y
					//TODO -- buug! item display disappears
					s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
					s.Screen.Buffer[s.Player.Y][s.Player.X].Set(item.Tile)
					s.Screen.Buffer[s.Player.Y][s.Player.X].Set(s.Player.Tile)
					s.Level.Items = append(s.Level.Items, item)
					s.Player.Items = append(s.Player.Items[:idx], s.Player.Items[idx+1:]...)
					s.Profile.SelectedItem = ""
					secondRefresh = true
					s.State = STATE_ITEMACTION
				}
		}
	}
	return secondRefresh
}