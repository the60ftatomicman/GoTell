package session

import (
	"strings"
)

func handleInputMenu(input string, s *Session) bool{
	previousState := s.State.Name
	lowerCaseInput := strings.ToLower(input)
	switch lowerCaseInput {
		case "a":{
			delta := -1
			if(s.Menu.GetSelection() == 0){
				s.Menu.ChangeClass(delta)
			}
			if(s.Menu.GetSelection() == 1){
				s.Menu.ChangeLevel(delta)
			}
		}
		case "d":{
			delta := 1
			if(s.Menu.GetSelection() == 0){
				s.Menu.ChangeClass(delta)
			}
			if(s.Menu.GetSelection() == 1){
				s.Menu.ChangeLevel(delta)
			}
		}
		case "w":{s.Menu.ChangeSelection(-1)}
		case "s":{s.Menu.ChangeSelection(1)}
		case "g":{
			s.Player = *s.Menu.Player
			s.State = STATE_MOVING
		}
		default: {}
	}
	return previousState != s.State.Name
}