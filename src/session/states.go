package session


type IState interface {
	IsInputValid(string) bool
}

type State struct {
	Name          string `default:""`
	validInputs []string
	waitForInput    bool
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
var KEY_QUIT = "Q"

var STATE_MOVING = State {
	Name: "moving",
	validInputs: KEY_DIRS,
	waitForInput: true,
	handleInput: handleInputMoving,
}
var STATE_INVENTORY = State {
	Name: "inventory",
	validInputs:[]string{"1","2","3","4","5","m"},
	waitForInput: true,
	handleInput: handleInputInventory,
}
var STATE_ITEM = State {
	Name: "item",
	validInputs:[]string{"u","c","d","i","m"},
	waitForInput: true,
	handleInput: handleInputItem,
}
var STATE_SPELL = State {
	Name: "castingspell",
	validInputs: append([]string{"x","m"},KEY_DIRS...),
	waitForInput: true,
	handleInput: handleInputSpell,
}
//Pass through states
var STATE_DEAD = State {
	Name: "dead",
	validInputs:[]string{"r"},
	waitForInput: true,
}
var STATE_GETITEM = State {
	Name: "pickingupitem",
	validInputs:[]string{},
	waitForInput: true, //TODOO -- need?
	handleInput: handleGetItem,
}