package simulater

import (
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/base/zone"
)

type Situation struct {
	Timer *Timer
	Ball  *Ball
	Side  Side
	Left  team.Ability
	Right team.Ability
}

func (this *Situation) Offender() team.Ability {
	if this.Side == Left {
		return this.Left
	} else {
		return this.Right
	}
}
func (this *Situation) Defender() team.Ability {
	if this.Side == Left {
		return this.Right
	} else {
		return this.Left
	}
}
func (this *Situation) Zone() zone.Name {
	return this.Ball.Zone(this.Side)
}