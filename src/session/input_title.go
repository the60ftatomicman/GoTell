package session

import (
	"strings"
)

func handleInputTitle(input string, s *Session) bool{
	previousState := s.State.Name
	lowerCaseInput := strings.ToLower(input)
	switch lowerCaseInput {
		case "s":{ s.State = STATE_STORY }
		case "b":{ s.State = STATE_MENU  }
		case "h":{}
		default: {}
	}
	return previousState != s.State.Name
}