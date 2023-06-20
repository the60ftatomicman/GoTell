package session

import (
	"bufio"
	"example/gotell/src/core"
	"example/gotell/src/core/screen"
	"example/gotell/src/core/tile"
	"example/gotell/src/object"
	"example/gotell/src/region"
	"fmt"
	"net"
	"strings"
)

type Session struct {
	Player     object.Player
	Screen     screen.Screen
	Header     region.Header
	Level      region.Level
	Profile    region.Profile
	Info       region.Info
	Popup      region.Popup
	Title      region.Splash
	Menu       region.Menu
	Story      []region.Splash
	currStory  int
	State      State
	Connection net.Conn
}

func (s *Session) Initialize(c *net.Conn) {
	// Set connection
	s.Connection = *c
	//S etup window
	core.HandleOutputToClient(s.Connection, 0, 0, core.Clear()+core.ResizeTerminal(screen.SCREEN_HEIGHT, screen.SCREEN_WIDTH))
	// Set all of our SPLASH Screens
	s.Title = region.Splash{FilePath: "/Users/andrew.garber/repo/funksi/GoTell/src/region/splash_screens/title.splash",}
	s.Title.Initialize([][]tile.Tile{})
	
	s.currStory = 0
	s.Story = []region.Splash{
		{FilePath: "/Users/andrew.garber/repo/funksi/GoTell/src/region/splash_screens/story_0.splash",},
		{FilePath: "/Users/andrew.garber/repo/funksi/GoTell/src/region/splash_screens/story_1.splash",},
	}
	for i := 0;i<len(s.Story);i++ {
		s.Story[i].Initialize([][]tile.Tile{})
	}
	
	// Set State
	s.State = STATE_TITLE
	

	s.Menu = region.Menu{}
	s.Menu.Initialize([][]tile.Tile{})
	//s.State = STATE_MOVING // -- DEBUG!
	
	//---------- Generate Player tile
	s.Player = object.GeneratePlayer()
	s.Screen = screen.Screen{
		Buffer: screen.BlankScreen(),
		Raw:    "",
	}
	// ---------- Generate Header region
	s.Header = region.Header{}
	s.Header.Initialize([][]tile.Tile{})
	// ---------- Generate Level region
	s.Level = region.Level{Player: &s.Player}
	s.Level.Initialize(s.Level.ReadDataFromFile())
	// ------------ Generate Profile region
	s.Profile = region.Profile{}
	s.Profile.Initialize(s.Profile.ReadDataFromPlayer(&s.Player))
	// ------------ Generate Info Region
	s.Info = region.Info{}
	s.Info.Initialize([][]tile.Tile{})
	// ------------ Generate Info Region
	s.Popup.Initialize([][]tile.Tile{})
}


//TODO -- have a map or function to determine which REGIONS we need to draw based on STATE.
// Maybe add to state?
func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	s.Screen.Compile(&s.Title) // -- actual prod
	//s.Screen.Compile(&s.Header,&s.Level, &s.Profile, &s.Info) // -- debugging gameplay
	s.Screen.Refresh()
	core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	//Begin Game loop
	for {
		netData, _    := bufio.NewReader(s.Connection).ReadByte()
		formattedData := strings.TrimSpace(string(netData))
		//AKA are we quitting
		if(handleGlobalStateSwitching(formattedData,s)){
			break
		}
		if handleInputStateSwitching(formattedData,s) || s.State.IsInputValid(formattedData){

			//This is a hack for getItem and THAT IS IT.
			if(s.State.handleInput(formattedData,s) ){
				s.State.handleInput(formattedData,s)
			}
			if(s.Popup.HasMessages()){
				s.State = STATE_POPUP
			}
			switch(s.State.Name){
				case STATE_POPUP.Name:{
					s.Popup.Refresh()
					s.Screen.Compile(&s.Profile, &s.Info,&s.Popup)
				}
				case STATE_TITLE.Name:{
					s.Screen.Compile(&s.Title)
				}
				case STATE_STORY.Name:{
					s.Screen.Compile(&s.Story[s.currStory])
				}
				case STATE_MENU.Name:{
					s.Screen.Compile(&s.Menu)
				}
				default:{
					//On map!
					s.Level.Refresh()
					s.Profile.Refresh()
					s.Info.Refresh()
					s.Screen.Compile(&s.Level,&s.Profile, &s.Info, &s.Header)
				}
			}
			s.Screen.Refresh()
		}
		core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	}
	s.Connection.Close()
}
