package simulate

import (
	"github.com/countingmars/fb/base"
)

type Situation struct {
	Timer *Timer
	Ball  *Ball
	Side  Side
	Left  base.TeamAbility
	Right base.TeamAbility
}

func (this *Situation) Offender() base.TeamAbility {
	if this.Side == Left {
		return this.Left
	} else {
		return this.Right
	}
}
func (this *Situation) Defender() base.TeamAbility {
	if this.Side == Left {
		return this.Right
	} else {
		return this.Left
	}
}