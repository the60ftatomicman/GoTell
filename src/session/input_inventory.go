package session

import "strconv"

func handleInputInventory(input string, s *Session){
	s.Info.Set("Currently [Invetory]: 1-3 (item), switch to (m)oving")
	idx,notInt := strconv.Atoi(input)
	if(notInt != nil){
		switch input {
			case "u":
				{
					s.Player.Items[idx].Interaction(&s.Player)
					s.Profile.Health = strconv.Itoa(s.Player.Stats.Health)
				}
			case "c":
				{

				}
		}
	}else{
		if(idx < len(s.Profile.Items)){
			s.Profile.SelectedItem = input
			s.Info.Set("["+s.Profile.Items[idx-1]+"] selected. (u)se? (c)onvert? (d)rop?")
		}
	}
	s.Profile.Refresh()
	s.Info.Refresh()
}