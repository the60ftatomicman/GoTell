package session

func handleInputInventory(input string, s *Session){
	s.Info.Set("Currently [Invetory]: 1-3 (item), switch to (m)oving")
	switch input {
		case "1":
			{

			}
		case "2":
			{

			}
	}
	s.Profile.Refresh()
	s.Info.Refresh()
}