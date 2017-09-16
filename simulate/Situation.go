package simulate

import (
	"github.com/countingmars/fb/core"
)

type Situation struct {
	Timer *Timer
	Ball *Ball
	Side Side
	Left core.TeamAbility
	Right core.TeamAbility
}

func (this *Situation) Offender() core.TeamAbility {
	if this.Side == Left {
		return this.Left
	} else {
		return this.Right
	}
}
func (this *Situation) Defender() core.TeamAbility {
	if this.Side == Left {
		return this.Right
	} else {
		return this.Left
	}
}