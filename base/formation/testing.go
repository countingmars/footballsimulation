package formation

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/player"
)

func TestPositionedPlayer(aPosition position.Position) *PositionedPlayer {
	aPlayer := player.TestPlayer(aPosition)
	aPlayer.Positions = position.Positions{aPosition}

	return &PositionedPlayer{
		Position: aPosition,
		Player:   aPlayer,
	}
}

func Test442() Formation {
	return Formation{
		TestPositionedPlayer(position.GK),
		TestPositionedPlayer(position.DC),
		TestPositionedPlayer(position.DC),
		TestPositionedPlayer(position.DL),
		TestPositionedPlayer(position.DR),
		TestPositionedPlayer(position.ML),
		TestPositionedPlayer(position.MR),
		TestPositionedPlayer(position.MC),
		TestPositionedPlayer(position.MC),
		TestPositionedPlayer(position.STC),
		TestPositionedPlayer(position.STC),
	}
}
