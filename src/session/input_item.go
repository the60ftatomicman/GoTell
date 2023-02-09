package session

import (
	"example/gotell/src/core"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"strconv"
)

func handleGetItem(input string, s *Session) bool{
	msg := "you are NOT on an item."
	if(len(s.Player.Items) < region.LINE_VAR_ITEM_COUNT){
		for idx,item := range s.Items {
			if (item.X == s.Player.X && item.Y == s.Player.Y){
				s.Player.Items = append(s.Player.Items,item)
				msg = "Picked up ["+item.Name+"]"
				//Remove player and item
				s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
				s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
				s.Screen.Buffer[s.Player.Y][s.Player.X].Set(s.Player.Tile)
				s.Items = append(s.Items[:idx], s.Items[idx+1:]...)
				if(tile.CheckAttributes(item.Tile,core.ATTR_EQUIPTABLE)){
					item.Interaction(&s.Player.Stats)
				}
			}
		}
	}else{
		msg = "Your inventory is FULL"
	}
	s.Info.Set(msg)
	s.State = STATE_MOVING
	return false
}

func handleInputItem(input string, s *Session) bool{
	idx,notInt         := strconv.Atoi(s.Profile.SelectedItem)
	if notInt == nil {
		//TODO -- error handling
		idx            = idx - 1
		item          := s.Player.Items[idx]
		s.Info.Set(
			"["+item.Name+"] selected.",
			"Press (u) to USE",
			"Press (c) to CONVERT",
			"Press (d) to DROP",
		)
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
						}
					}
				}
			case "c":
				{
					s.Player.Stats.ChangeXP(item.ConversionPoints)
					s.Player.Items = append(s.Player.Items[:idx], s.Player.Items[idx+1:]...)
					s.Profile.SelectedItem = ""
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
					s.Items = append(s.Items, item)
					s.Player.Items = append(s.Player.Items[:idx], s.Player.Items[idx+1:]...)
					s.Profile.SelectedItem = ""
				}
		}
		s.Info.Refresh()
	}
	return false
}