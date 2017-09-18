package formation

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/player"
)

func TestRole(aPosition position.Position) *Role {
	aPlayer := player.TestPlayer(aPosition)
	aPlayer.Positions = position.Positions{aPosition}

	return &Role{
		Position: aPosition,
		Player:   aPlayer,
	}
}

func Test442() Formation {
	return Formation{
		TestRole(position.GK),
		TestRole(position.DC),
		TestRole(position.DC),
		TestRole(position.DL),
		TestRole(position.DR),
		TestRole(position.ML),
		TestRole(position.MR),
		TestRole(position.MC),
		TestRole(position.MC),
		TestRole(position.STC),
		TestRole(position.STC),
	}
}
