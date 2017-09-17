package simulater

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/base/zone"
)

type Situation struct {
	Timer *Timer
	Ball  *Ball
	Side  Side
	Left  zone.Zones
	Right zone.Zones
}

func (this *Situation) Offender() zone.Zones {
	if this.Side == Left {
		return this.Left
	} else {
		return this.Right
	}
}
func (this *Situation) Defender() zone.Zones {
	if this.Side == Left {
		return this.Right
	} else {
		return this.Left
	}
}
func (this *Situation) ZoneName() base.Name {
	return this.Ball.ZoneName(this.Side)
}