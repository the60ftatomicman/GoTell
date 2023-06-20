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
			if(s.Menu.CursorIdx == 0){
				s.Menu.ChangeClass(delta)
			}
			if(s.Menu.CursorIdx == 1){
				s.Menu.ChangeLevel(delta)
			}
		}
		case "w":{s.Menu.MoveCursor(-1)}
		case "d":{
			delta := 1
			if(s.Menu.CursorIdx == 0){
				s.Menu.ChangeClass(delta)
			}
			if(s.Menu.CursorIdx == 1){
				s.Menu.ChangeLevel(delta)
			}
		}
		case "s":{s.Menu.MoveCursor(1)}
		case "g":{
			s.Player = *s.Menu.Player
			s.State = STATE_MOVING
		}
		default: {}
	}
	return previousState != s.State.Name
}