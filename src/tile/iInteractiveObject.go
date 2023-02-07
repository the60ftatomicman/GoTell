package tile

//Seems uselesss but want to get working
type IInteractiveObject interface {
	GetBufferData() (int,int,string,Tile)
	Interaction(p *Player) bool
	Convert(p *Player)
}
