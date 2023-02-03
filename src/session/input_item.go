package session

import "strconv"

func handleInputItem(input string, s *Session){
	s.Info.Set("Currently [Invetory]: 1-3 (item), switch to (m)oving")
	idx,notInt := strconv.ParseInt(input,10,0)
	if(notInt != nil){
		switch input {
			case "u":
				{
					s.Player.Items[idx].Interaction(&s.Player)
				}
			case "c":
				{

				}
		}
	}else{
		s.Profile.SelectedItem = input
		s.Info.Set("["+s.Profile.Items[idx]+"] selected. (u)se? (c)onvert? (d)rop?")
	}
	s.Profile.Refresh()
	s.Info.Refresh()
}