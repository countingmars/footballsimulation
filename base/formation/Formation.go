package formation

import (
	"github.com/countingmars/fb/base/position"
)

var F442 = Formation {
	NewPositionedPlayer(position.GK),
	NewPositionedPlayer(position.DL),
	NewPositionedPlayer(position.DR),
	NewPositionedPlayer(position.DC),
	NewPositionedPlayer(position.DC),
	NewPositionedPlayer(position.ML),
	NewPositionedPlayer(position.MR),
	NewPositionedPlayer(position.MC),
	NewPositionedPlayer(position.MC),
	NewPositionedPlayer(position.STC),
	NewPositionedPlayer(position.STC),
}

type Formation []*PositionedPlayer

func (this Formation) Set(point int) {
	for _, positioned := range this {
		positioned.Player.Ability.Set(point)
	}
}

func (this Formation) Clone() Formation {
	clone := Formation{}
	for _, positioned := range this {
		clone = append(clone, positioned.Clone())
	}
	return clone
}
