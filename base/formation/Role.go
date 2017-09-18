package formation

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/player"
)

type Role struct {
	Position position.Position `json: "position"`
	Player *player.Player `json: "player"`
}

func (this Role) Clone() *Role {
	player := this.Player.Clone()
	return &Role{
		Position: this.Position,
		Player: player,
	}
}
func NewRole(aPosition position.Position) *Role {
	return &Role{Position: aPosition}
}
