package session


func handleInputPopup(input string,s *Session) bool{
	switch input {
		case "y":
			{
				s.Popup.ClearMessages()
				s.State = STATE_MOVING
			}
	}
	return false
}