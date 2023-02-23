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

}

//TODO -- theseOUGHT to be a level region function
func (s *Session) placeObject(interObj tile.IInteractiveObject)	{
		objY,objX,objName,objTile := interObj.GetBufferData()
		intendedType := s.Screen.Buffer[objY][objX].Get()
		if (tile.CheckAttributes(intendedType,core.ATTR_SOLID)){
			fmt.Println("ERROR placing item ["+objName+"] at location ["+strconv.Itoa(objY)+"]["+strconv.Itoa(objX)+"] do to ["+intendedType.Name+"] tile which is solid")
		}
		s.Screen.Buffer[objY][objX].Pop()
		s.Screen.Buffer[objY][objX].Pop()
		s.Screen.Set(objTile, objY,objX)
		s.Screen.Set(tile.FOG, objY,objX)
}

//TODO -- get rid of this somehow? Keep sessions out of level logic
func (s *Session) initializeObjects() {
	//--Enemies
	for _,enemy := range s.Level.Enemies {
		s.placeObject(&enemy)
	}
	//--Items
	for _,item := range s.Level.Items {
		s.placeObject(&item)
	}
	//-- Remove initial Fog around player
	for r:=s.Player.Stats.Vision * -1;r<s.Player.Stats.Vision ;r++{
		for c:=s.Player.Stats.Vision * -1;c<s.Player.Stats.Vision ;c++{
			visionR := r+s.Player.Y
			visionC := c+s.Player.X
			inColumn := visionC >=0 && visionC < region.MAP_LEFT+region.MAP_COLUMNS
			inRow    := visionR >=0 && visionR < region.MAP_TOP+region.MAP_LINES
			if(inColumn && inRow){
				cell := s.Screen.Buffer[visionR][visionC].Get()
				if(cell.Name == tile.FOG.Name){
					s.Screen.Buffer[visionR][visionC].Pop()
				}
			}
		}
	}
	//-- Now place player
	s.Screen.Set(s.Player.Tile, s.Player.Y, s.Player.X)
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	s.Screen.Compile(&s.Header,&s.Level, &s.Profile, &s.Info)
	s.initializeObjects()
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
			if(s.State.handleInput(formattedData,s)){
				//This is a hack for getItem and THAT IS IT.
				s.State.handleInput(formattedData,s)
			}
			//Refresh our dynamic regions
			s.Profile.Player = &s.Player
			s.Profile.Refresh()
			s.Info.Refresh()
			//Refresh the full screen.
			s.Screen.Compile(&s.Profile, &s.Info)
			s.Screen.Refresh()
		}
		core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	}
	s.Connection.Close()
}
