package simulater

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/simulater/dice"
	"fmt"
)

type Ball struct {
	x int
	y int
}


func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
}
func (this *Ball) Forward(side Side) base.Zone {
	this.y = dice.Throw(3)
	if side == Left {
		this.x++
	} else {
		this.x--
	}
	return this.Zone(side)
}
func (this *Ball) Zone(side Side) base.Zone {
	key := fmt.Sprintf("%d%d%s", this.x, this.y, side)
	return field[key]
}


func (this *Ball) CanFinish(side Side) bool {
	return this.Zone(side) == base.Zones.GK
}

var field = map[string]base.Zone {
	"00left":   base.Zones.DR,
	"01left":   base.Zones.DC,
	"02left":   base.Zones.DL,
	"10left":   base.Zones.MR,
	"11left":   base.Zones.MC,
	"12left":   base.Zones.ML,
	"20left":   base.Zones.FR,
	"21left":   base.Zones.FC,
	"22left":   base.Zones.FL,
	"30left":   base.Zones.GK,
	"31left":   base.Zones.GK,
	"32left":   base.Zones.GK,
	"-10right": base.Zones.GK,
	"-11right": base.Zones.GK,
	"-12right": base.Zones.GK,
	"00right":  base.Zones.FR,
	"01right":  base.Zones.FC,
	"02right":  base.Zones.FL,
	"10right":  base.Zones.MR,
	"11right":  base.Zones.MC,
	"12right":  base.Zones.ML,
	"20right":  base.Zones.DR,
	"21right":  base.Zones.DC,
	"22right":  base.Zones.DL,
}