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
	Level      region.Level
	Profile    region.Profile
	Info       region.Info
	State      string //TODO -- ENUM!
	Enemies    []tile.Enemy
	Items      []tile.Item
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
	// ------------ Generate Enemies
	s.Enemies = tile.GenerateEnemiesFromFile()
	// ------------ Generate Items
	s.Items = tile.GenerateItemsFromFile()
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
		if (tile.CheckAttributes(intendedType,core.ATTR_FOREGROUND)){
			s.Screen.Buffer[objY][objX].Pop()
			s.Screen.Set(objTile, objY,objX)
			s.Screen.Set(tile.FOG, objY,objX)
		}else{
			s.Screen.Set(objTile, objY,objX)
		}
}


func (s *Session) initializeObjects() {
	//--Enemies
	for _,enemy := range s.Enemies {
		s.placeObject(&enemy)
	}
	//--Items
	for _,item := range s.Items {
		s.placeObject(&item)
	}
	s.Screen.Set(s.Player.Tile, s.Player.Y, s.Player.X)
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	s.Screen.Compile(&s.Level, &s.Profile, &s.Info)
	s.initializeObjects()
	s.Screen.Refresh()
	core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	for {
		netData, _    := bufio.NewReader(s.Connection).ReadByte()
		formattedData := strings.TrimSpace(string(netData))
		if hanleInputStateSwitching(formattedData, s) {
			// AKA Quit
			break
		} else {
			switch s.State{
				case STATE_MOVING:{
					// TODO -- we only need the session
					handleInputMoving(formattedData, &s.Player, s)
				}
				case STATE_INVENTORY:{
					handleInputInventory(formattedData, s)
				}
			}
			s.Profile.Player = &s.Player
			s.Profile.Refresh()
			s.Screen.Compile(&s.Profile, &s.Info)
			s.Screen.Refresh()
			core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
		}
	}
	s.Connection.Close()
}
