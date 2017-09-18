package zone

import (
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/player"
)

func TestZone(name name.Name, array ...*Entry) Zone {
	entries := Entries{}
	entries = append(entries, array...)
	return Zone{
		Name:    name,
		Entries: entries,
	}
}
func TestEntry(aPosition position.Position) *Entry {
	aPlayer := player.TestPlayer(aPosition)
	aPlayer.Positions = position.Positions{aPosition}

	return &Entry{
		Position: aPosition,
		Player:   aPlayer,
		Stats: Stats{},
	}
}
