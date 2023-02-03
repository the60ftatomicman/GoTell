package tile

type iInteractiveObject interface {
	Interaction(p *Player) bool
	Convert(p *Player)
}
