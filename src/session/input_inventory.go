package session

import (
	"example/gotell/src/core"
	"example/gotell/src/tile"
	"strconv"
)

func handleInputInventory(input string, s *Session){
	//TODO -- make a "menus" section similar to TILES and what not
	s.Info.Set(
		"Currently viewing [Invetory]",
		"Press [1-5] to select an item",
		"Press [M] to return to moving",
	)
	idx,notInt := strconv.Atoi(input)
	if(notInt != nil && s.Profile.SelectedItem != ""){
		//TODO - make method in profile!
		idx,_ = strconv.Atoi(s.Profile.SelectedItem)
		idx = idx - 1
		item := s.Player.Items[idx]
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
	}else{
		if(idx > 0 && idx <= len(s.Profile.Player.Items)){
			idx -= 1
			s.Profile.SelectedItem = input
			item := s.Profile.Player.Items[idx]
			if(!tile.CheckAttributes(item.Tile,core.ATTR_EQUIPTABLE)){
				s.Info.Set(
					"["+item.Name+"] selected. Press (b) to go BACK",
					"Press (u) to USE",
					"Press (c) to CONVERT",
					"Press (d) to DROP")
			}else{
				s.Info.Set(
					"["+item.Name+"] selected. Press (b) to go BACK",
					"Press (c) to CONVERT",
					"Press (d) to DROP")	
			}
		}else{
			s.Info.Set(
				"Currently viewing [Invetory]",
				"OOPS! ["+strconv.Itoa(idx)+"] is just not valid input!",
				"Press [1-5] to select an item",
				"Press [M] to return to moving",
			)	
		}
	}
	s.Info.Refresh()
}