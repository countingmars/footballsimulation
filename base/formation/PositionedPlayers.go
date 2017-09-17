package formation


type PositionedPlayers []*PositionedPlayer

func (this PositionedPlayers) Clone() PositionedPlayers {
	clone := PositionedPlayers{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}