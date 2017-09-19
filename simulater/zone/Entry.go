package zone

import (
	"github.com/countingmars/fb/base/player"
	"github.com/countingmars/fb/base/position"
)

type Entry struct {
	Player *player.Player
	Position position.Position
	Stats Stats
}

func (this Entry) Clone() *Entry {
	return &Entry{
		Position: this.Position,
		Player: this.Player.Clone(),
		Stats: this.Stats.Clone(),
	}
}
