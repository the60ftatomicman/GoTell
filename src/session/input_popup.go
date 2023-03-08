package session



func handleInputPopup(input string,s *Session) bool{
	switch input {
		case "y":
			{
				s.State = STATE_MOVING
			}
		case "n":
			{
				s.State = STATE_MOVING
			}
	}
	return false
}