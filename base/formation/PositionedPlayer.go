package formation

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/player"
)

type PositionedPlayer struct {
	Position position.Position `json: "position"`
	Player *player.Player `json: "player"`
}

func (this PositionedPlayer) Clone() *PositionedPlayer {
	player := this.Player.Clone()
	return &PositionedPlayer{
		Position: this.Position,
		Player: player,
	}
}
func NewPositionedPlayer(aPosition position.Position) *PositionedPlayer {
	return &PositionedPlayer{Position: aPosition}
}
