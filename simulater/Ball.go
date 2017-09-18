package simulater

import (
	"fmt"
	"github.com/countingmars/fb/simulater/dice"
	"github.com/countingmars/fb/base/name"
)

type Ball struct {
	x int
	y int
}


func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
}
func (this *Ball) Forward(side Side) name.Name {
	this.y = dice.Throw(3)
	if side == Left {
		this.x++
	} else {
		this.x--
	}
	return this.ZoneName(side)
}
func (this *Ball) ZoneName(side Side) name.Name {
	key := fmt.Sprintf("%d%d%s", this.x, this.y, side)
	return field[key]
}


func (this *Ball) CanFinish(side Side) bool {
	return this.ZoneName(side) == name.GK
}

var field = map[string]name.Name{
	"00left":   name.DR,
	"01left":   name.DC,
	"02left":   name.DL,
	"10left":   name.MR,
	"11left":   name.MC,
	"12left":   name.ML,
	"20left":   name.FR,
	"21left":   name.FC,
	"22left":   name.FL,
	"30left":   name.GK,
	"31left":   name.GK,
	"32left":   name.GK,
	"-10right": name.GK,
	"-11right": name.GK,
	"-12right": name.GK,
	"00right":  name.FR,
	"01right":  name.FC,
	"02right":  name.FL,
	"10right":  name.MR,
	"11right":  name.MC,
	"12right":  name.ML,
	"20right":  name.DR,
	"21right":  name.DC,
	"22right":  name.DL,
}