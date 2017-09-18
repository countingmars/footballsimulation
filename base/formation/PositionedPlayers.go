package formation

import "encoding/json"

type PositionedPlayers []*PositionedPlayer

func (this PositionedPlayers) Clone() PositionedPlayers {
	clone := PositionedPlayers{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}
func (this PositionedPlayers) Json() string {
	out, _ := json.Marshal(this)
	return string(out)
}