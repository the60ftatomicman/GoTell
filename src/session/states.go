package session


type IState interface {
	IsInputValid(string) bool
}

type State struct {
	Name          string `default:""`
	validInputs []string
	handleInput     func(input string,s *Session) bool
}

func (s *State) IsInputValid(input string) bool{
    for _, b := range s.validInputs {
        if b == input {
            return true
        }
    }
    return false
}

var KEY_DIRS = []string{"w","a","s","d"}
var KEY_MENU = []string{"w","a","s","d","g"}
var KEY_QUIT = "Q"

var STATE_TITLE = State {
	Name: "title",
	validInputs: []string{"b","B","h","H","s","S","q","Q"},
	handleInput: handleInputTitle,
}
var STATE_STORY = State {
	Name: "story",
	validInputs: []string{"a","A","d","D","q"},
	handleInput: handleInputStory,
}
var STATE_MENU = State {
	Name: "menu",
	validInputs: KEY_MENU,
	handleInput: handleInputMenu,
}
var STATE_MOVING = State {
	Name: "moving",
	validInputs: KEY_DIRS,
	handleInput: handleInputMoving,
}
var STATE_INVENTORY = State {
	Name: "inventory",
	validInputs:[]string{"1","2","3","4","5","m"},
	handleInput: handleInputInventory,
}
var STATE_ITEM = State {
	Name: "item",
	validInputs:[]string{"u","c","d","i","m"},
	handleInput: handleInputItem,
}
var STATE_SPELL = State {
	Name: "castingspell",
	validInputs: append([]string{"x","m"},KEY_DIRS...),
	handleInput: handleInputSpell,
}
var STATE_POPUP = State {
	Name: "popup",
	validInputs: []string{"y","n"},
	handleInput: handleInputPopup,
}
//Pass through states
var STATE_DEAD = State {
	Name: "dead",
	validInputs:[]string{"r"},
}
var STATE_GETITEM = State {
	Name: "pickingupitem",
	validInputs:[]string{},
	handleInput: handleGetItem,
}
var STATE_ITEMACTION= State {
	Name: "itemaction",
	validInputs:[]string{},
	handleInput: handleItemAction,
}