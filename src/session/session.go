package session

import (
	"bufio"
	"example/gotell/src/core"
	"example/gotell/src/screen"
	"example/gotell/src/screen/region"
	"example/gotell/src/tile"
	"fmt"
	"net"
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
	// ---------- Generate Level region
	s.Level = region.Level{}
	s.Level.Initialize(s.Level.ReadDataFromFile())
	// ------------ Generate Profile region
	s.Profile = region.Profile{}
	s.Profile.Initialize(s.Profile.ReadDataFromFile())
	// ------------ Generate Info Region
	s.Info = region.Info{}
	s.Info.Initialize([][]tile.Tile{})

}

func (s *Session) initializeObjects() {
	//--Enemies
	for idx,_ := range s.Enemies {
		s.Screen.Set(s.Enemies[idx].Tile, s.Enemies[idx].X,s.Enemies[idx].Y)
	}
	//--Player
	//if(playerMoved){
	//	s.Screen.Set(s.Player.UnderTile, s.Player.PrvX, s.Player.PrvY)
	//	s.Player.UnderTile = s.Screen.Buffer[s.Player.Y][s.Player.X].Get()
	//}
	s.Screen.Set(s.Player.Tile, s.Player.X, s.Player.Y)
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	s.initializeObjects()
	s.Screen.Compile(&s.Level, &s.Profile, &s.Info)
	core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
	for {
		netData, _    := bufio.NewReader(s.Connection).ReadByte()
		formattedData := strings.TrimSpace(string(netData))
		if hanleInputStateSwitching(formattedData, s, &s.Info) {
			// AKA Quit
			break
		} else {
			switch s.State{
				case STATE_MOVING:{
					// TODO -- we only need the buffer....
					handleInputMoving(formattedData, &s.Player, s)
				}
				case STATE_INVENTORY:{
					handleInputInventory(formattedData, &s.Profile, &s.Info)
				}
			}
			s.Screen.Compile(&s.Level, &s.Profile, &s.Info)
			core.HandleOutputToClient(s.Connection, 0, region.INFO_TOP+region.INFO_LINES+1, s.Screen.Get())
		}
	}
	s.Connection.Close()
}
