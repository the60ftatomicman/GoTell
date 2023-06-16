package session

import (
	"strings"
)

func handleInputStory(input string, s *Session) bool{
	previousState := s.State.Name
	lowerCaseInput := strings.ToLower(input)
	switch lowerCaseInput {
		case "a":{
			if(s.currStory > 0){
				s.currStory--
			}
		}
		case "d":
			{
				s.currStory += 1
				if(s.currStory >= len(s.Story)){
					s.currStory = 0
					s.State = STATE_MOVING
				}
			}
		case "q":{s.State = STATE_MOVING}
		default: {}
	}
	return previousState != s.State.Name
}