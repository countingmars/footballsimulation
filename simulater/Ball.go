package simulater

import (
	"fmt"
	"github.com/countingmars/fb/simulater/dice"
	"github.com/countingmars/fb/base/zone"
)

type Ball struct {
	x int
	y int
}


func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
}
func (this *Ball) Forward(side Side) zone.Name {
	this.y = dice.Throw(3)
	if side == Left {
		this.x++
	} else {
		this.x--
	}
	return this.Zone(side)
}
func (this *Ball) Zone(side Side) zone.Name {
	key := fmt.Sprintf("%d%d%s", this.x, this.y, side)
	return field[key]
}


func (this *Ball) CanFinish(side Side) bool {
	return this.Zone(side) == zone.Names.GK
}

var field = map[string]zone.Name{
	"00left":   zone.Names.DR,
	"01left":   zone.Names.DC,
	"02left":   zone.Names.DL,
	"10left":   zone.Names.MR,
	"11left":   zone.Names.MC,
	"12left":   zone.Names.ML,
	"20left":   zone.Names.FR,
	"21left":   zone.Names.FC,
	"22left":   zone.Names.FL,
	"30left":   zone.Names.GK,
	"31left":   zone.Names.GK,
	"32left":   zone.Names.GK,
	"-10right": zone.Names.GK,
	"-11right": zone.Names.GK,
	"-12right": zone.Names.GK,
	"00right":  zone.Names.FR,
	"01right":  zone.Names.FC,
	"02right":  zone.Names.FL,
	"10right":  zone.Names.MR,
	"11right":  zone.Names.MC,
	"12right":  zone.Names.ML,
	"20right":  zone.Names.DR,
	"21right":  zone.Names.DC,
	"22right":  zone.Names.DL,
}