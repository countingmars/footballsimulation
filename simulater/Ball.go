package simulater

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/simulater/dice"
)

type Ball struct {
	x int
	y int
	possessor Side
}

var field = map[Ball]base.Zone{
	{0, 0, Left}:   base.Zones.DR,
	{0, 1, Left}:   base.Zones.DC,
	{0, 2, Left}:   base.Zones.DL,
	{1, 0, Left}:   base.Zones.MR,
	{1, 1, Left}:   base.Zones.MC,
	{1, 2, Left}:   base.Zones.ML,
	{2, 0, Left}:   base.Zones.FR,
	{2, 1, Left}:   base.Zones.FC,
	{2, 2, Left}:   base.Zones.FL,
	{3, 0, Left}:   base.Zones.GK,
	{3, 1, Left}:   base.Zones.GK,
	{3, 2, Left}:   base.Zones.GK,
	{-1, 0, Right}: base.Zones.GK,
	{-1, 1, Right}: base.Zones.GK,
	{-1, 2, Right}: base.Zones.GK,
	{0, 0, Right}:  base.Zones.FR,
	{0, 1, Right}:  base.Zones.FC,
	{0, 2, Right}:  base.Zones.FL,
	{1, 0, Right}:  base.Zones.MR,
	{1, 1, Right}:  base.Zones.MC,
	{1, 2, Right}:  base.Zones.ML,
	{2, 0, Right}:  base.Zones.DR,
	{2, 1, Right}:  base.Zones.DC,
	{2, 2, Right}:  base.Zones.DL,

}
func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
	this.possessor = Left
}
func (this *Ball) Forward() base.Zone {
	this.y = dice.Throw(3)
	if this.possessor == Left {
		this.x++
	} else {
		this.x--
	}
	return this.Zone()
}
func (this *Ball) Zone() base.Zone {
	if this.possessor == Left {
		return field[*this]
	} else {
		return field[*this]
	}
}


func (this *Ball) CanFinish() bool {
	return this.Zone() == base.Zones.GK
}
