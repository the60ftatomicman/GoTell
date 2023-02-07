package tile

//Seems uselesss but want to get working
type IInteractiveObject interface {
	GetBufferData() (int,int,string,Tile)
	Interaction(s *Stats) bool
	Convert(s *Stats)
}
