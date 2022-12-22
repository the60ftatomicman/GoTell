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
	Connection net.Conn
}

func (s *Session) Initialize(c *net.Conn) {
	//Set connection
	s.Connection = *c
	//Setup window
	core.HandleOutputToClient(s.Connection, 0, 0, core.Clear()+core.ResizeTerminal(screen.SCREEN_HEIGHT, screen.SCREEN_WIDTH))
	//Generate Player
	s.Player = tile.GeneratePlayer()
	s.Screen = screen.Screen{
		Buffer: screen.BlankScreen(),
		Raw:    "",
	}
	/// ----------
	s.Level = region.Level{
		Name: "tutorial",
	}
	s.Level.Initialize([][]tile.Tile{
		{tile.BLANK},
		{tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK, tile.BLANK},
		{tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL, tile.WALL},
	})
	// --------------
	s.Profile = region.Profile{
		Name: "Hero",
	}
	s.Profile.Initialize([][]tile.Tile{})
}

func (s *Session) Handle() {
	fmt.Printf("Serving %s\n", s.Connection.RemoteAddr().String())
	//m.loadMap(&s)
	//weredude := tile.GenerateEnemy()
	//s.Screen.Set(weredude.Tile, weredude.X, weredude.Y)
	for {
		netData, _ := bufio.NewReader(s.Connection).ReadByte()

		quit := handleInput(strings.TrimSpace(string(netData)), &s.Player, &s.Screen)
		if quit {
			break
		} else {
			s.Screen.Compile(&s.Level, &s.Profile)
			core.HandleOutputToClient(s.Connection, 0, region.MAP_LINES+1, s.Screen.Get())
		}
	}
	s.Connection.Close()
}
