package zone

import (
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/base/formation"
)

func TestZone(name name.Name, array ...*formation.PositionedPlayer) Zone {
	players := formation.PositionedPlayers{}
	players = append(players, array...)
	return Zone{
		Name: name,
		PositionedPlayers: players,
	}
}
