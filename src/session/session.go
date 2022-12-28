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

const STATE_MOVING = "moving"
const STATE_INVENTORY = "inventory"
const STATE_ITEM = "item"


type Session struct {
	Player     tile.Player
	Screen     screen.Screen
	Level      region.Level
	Profile    region.Profile
	Info       region.Info
	State      string //TODO -- ENUM!
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
	/// ---------- Generate Level region
	s.Level = region.Level{}
	s.Level.Initialize(s.Level.ReadDataFromFile())
	// ------------ Generate Profile region
	s.Profile = region.Profile{}
	s.Profile.Initialize(s.Profile.ReadDataFromFile())
	// ------------ Generate Info Region
	s.Info = region.Info{}
	s.Info.Initialize([][]tile.Tile{})
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	//weredude := tile.GenerateEnemy()
	//s.Screen.Set(weredude.Tile, weredude.X, weredude.Y)
	for {
		netData, _    := bufio.NewReader(s.Connection).ReadByte()
		formattedData := strings.TrimSpace(string(netData))
		
		if hanleInputStateSwitching(formattedData, s, &s.Info) {
			// AKA Quit
			break
		} else {
			switch s.State{
				case STATE_MOVING:{
					handleInputMoving(formattedData, &s.Player, &s.Screen)
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
