package object

import "example/gotell/src/core/tile"

//Seems uselesss but want to get working
type IInteractiveObject interface {
	GetBufferData() (int,int,string,tile.Tile)
	Interaction(s *Stats) bool
	Convert(s *Stats)
}
