package session

import (
	"strings"
)

func handleInputMenu(input string, s *Session) bool{
	previousState := s.State.Name
	lowerCaseInput := strings.ToLower(input)
	switch lowerCaseInput {
		case "a":{}
		case "w":{}
		case "d":{}
		case "s":{s.State = STATE_MOVING}
		case "b":{}
		default: {}
	}
	return previousState != s.State.Name
}