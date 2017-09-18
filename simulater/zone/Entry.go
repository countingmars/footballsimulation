package zone

import (
	"github.com/countingmars/fb/base/player"
	"github.com/countingmars/fb/base/position"
	"encoding/json"
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

type Entries []*Entry

func (this Entries) Clone() Entries {
	clone := Entries{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}
func (this Entries) Json() string {
	out, _ := json.Marshal(this)
	return string(out)
}
