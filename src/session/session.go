package session

import (
	"bufio"
	"example/gotell/src/core"
	"example/gotell/src/screen"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Session struct {
	Player     tile.Player
	Screen     screen.Screen
	Header     region.Header
	Level      region.Level
	Profile    region.Profile
	Info       region.Info
	Popup      region.Popup
	State      State
	Connection net.Conn
}

func (s *Session) Initialize(c *net.Conn) {
	// Set connection
	s.Connection = *c
	//S etup window
	core.HandleOutputToClient(s.Connection, 0, 0, core.Clear()+core.ResizeTerminal(screen.SCREEN_HEIGHT, screen.SCREEN_WIDTH))
	// Set Staet
	s.State = STATE_MOVING
	//---------- Generate Player tile
	s.Player = tile.GeneratePlayer()
	s.Screen = screen.Screen{
		Buffer: screen.BlankScreen(),
		Raw:    "",
	}
	// ---------- Generate Header region
	s.Header = region.Header{}
	s.Header.Initialize([][]tile.Tile{})
	// ---------- Generate Level region
	s.Level = region.Level{}
	s.Level.Initialize(s.Level.ReadDataFromFile())
	// ------------ Generate Profile region
	s.Profile = region.Profile{}
	s.Profile.Initialize(s.Profile.ReadDataFromPlayer(&s.Player))
	// ------------ Generate Info Region
	s.Info = region.Info{}
	s.Info.Initialize([][]tile.Tile{})
	// ------------ Generate Info Region
	s.Popup = region.Popup{}
	s.Popup.Initialize([][]tile.Tile{})
	s.Popup.Set("Some dumb message")
	s.Popup.Refresh()
}

//TODO -- these OUGHT to be a level region function
func (s *Session) placeObject(interObj tile.IInteractiveObject)	{
		objY,objX,objName,objTile := interObj.GetBufferData()
		intendedType := s.Screen.Buffer[objY][objX].Get()
		if (tile.CheckAttributes(intendedType,core.ATTR_SOLID)){
			fmt.Println("ERROR placing item ["+objName+"] at location ["+strconv.Itoa(objY)+"]["+strconv.Itoa(objX)+"] do to ["+intendedType.Name+"] tile which is solid")
		}
		s.Screen.Buffer[objY][objX].Pop()
		s.Screen.Buffer[objY][objX].Pop()
		s.Screen.Set(objTile, objY,objX)
}

//TODO -- Put this WHOOOOLE thing into the level logic. Leevl should just have a pointer to the player obj.
func (s *Session) initializeObjects() {
	//--Enemies
	for _,enemy := range s.Level.Enemies {
		s.placeObject(&enemy)
	}
	//--Items
	for _,item := range s.Level.Items {
		s.placeObject(&item)
	}
	//-- Now place player
	s.Screen.Buffer[s.Player.Y][s.Player.X].Pop()
	s.Screen.Set(s.Player.Tile, s.Player.Y, s.Player.X)

	//-- Set FOG and other mask . Do not set fog around player
	for r:=region.MAP_TOP;r < region.MAP_TOP+region.MAP_LINES;r++ {
		for c:=region.MAP_LEFT;c < region.MAP_LEFT+region.MAP_COLUMNS;c++ {
			inRow := r >= s.Player.Y - s.Player.Stats.Vision && r <= s.Player.Y + s.Player.Stats.Vision
			inCol := c >= s.Player.X - s.Player.Stats.Vision && c <= s.Player.X + s.Player.Stats.Vision
			if(!inRow || !inCol){
				s.Screen.Buffer[r][c].Set(tile.FOG)
			}
		}
	}
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	s.initializeObjects()
	s.Screen.Compile(&s.Header,&s.Level, &s.Profile, &s.Info)
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
			formerState := s.State.Name // Hack for popups

			//This is a hack for getItem and THAT IS IT.
			if(s.State.handleInput(formattedData,s) ){
				s.State.handleInput(formattedData,s)
			}
			if(s.State.Name == STATE_POPUP.Name){
				s.Popup.Refresh()
				s.Screen.Compile(&s.Profile, &s.Info,&s.Popup)
			}else{
				//Refresh our dynamic regions
				s.Profile.Player = &s.Player
				s.Profile.Refresh()
				s.Info.Refresh()
				//Refresh the full screen.
				if(formerState == STATE_POPUP.Name){
					s.Screen.Compile(&s.Level,&s.Profile, &s.Info)
				}else{
					s.Screen.Compile(&s.Profile, &s.Info)
				}
			}
			s.Screen.Refresh()
		}
		core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	}
	s.Connection.Close()
}
