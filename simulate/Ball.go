package simulate

import (
	"github.com/countingmars/fb/core"
	"github.com/countingmars/fb/simulate/dice"
)

type Ball struct {
	x int
	y int
	possessor Side
}

var field = map[Ball]core.Zone{
	{0, 0, Left}: core.Zones.DR,
	{0, 1, Left}: core.Zones.DC,
	{0, 2, Left}: core.Zones.DL,
	{1, 0, Left}: core.Zones.MR,
	{1, 1, Left}: core.Zones.MC,
	{1, 2, Left}: core.Zones.ML,
	{2, 0, Left}: core.Zones.FR,
	{2, 1, Left}: core.Zones.FC,
	{2, 2, Left}: core.Zones.FL,
	{3, 0, Left}: core.Zones.GK,
	{3, 1, Left}: core.Zones.GK,
	{3, 2, Left}: core.Zones.GK,
	{-1, 0, Right}: core.Zones.GK,
	{-1, 1, Right}: core.Zones.GK,
	{-1, 2, Right}: core.Zones.GK,
	{0, 0, Right}: core.Zones.FR,
	{0, 1, Right}: core.Zones.FC,
	{0, 2, Right}: core.Zones.FL,
	{1, 0, Right}: core.Zones.MR,
	{1, 1, Right}: core.Zones.MC,
	{1, 2, Right}: core.Zones.ML,
	{2, 0, Right}: core.Zones.DR,
	{2, 1, Right}: core.Zones.DC,
	{2, 2, Right}: core.Zones.DL,

}
func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
	this.possessor = Left
}
func (this *Ball) Forward() core.Zone {
	this.y = dice.Throw(3)
	if this.possessor == Left {
		this.x++
	} else {
		this.x--
	}
	return this.Zone()
}
func (this *Ball) Zone() core.Zone {
	if this.possessor == Left {
		return field[*this]
	} else {
		return field[*this]
	}
}


func (this *Ball) CanFinish() bool {
	return this.Zone() == core.Zones.GK
}
