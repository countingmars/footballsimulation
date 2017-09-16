package simulator

import (
	. "github.com/countingmars/fb/foundation"
)

type Ball struct {
	x int
	y int
	possessor Side
}
var field = map[Ball]Zone{
	{0, 0, Left}: Zones.DR,
	{0, 1, Left}: Zones.DC,
	{0, 2, Left}: Zones.DL,
	{1, 0, Left}: Zones.MR,
	{1, 1, Left}: Zones.MC,
	{1, 2, Left}: Zones.ML,
	{2, 0, Left}: Zones.FR,
	{2, 1, Left}: Zones.FC,
	{2, 2, Left}: Zones.FL,
	{3, 0, Left}: Zones.GK,
	{3, 1, Left}: Zones.GK,
	{3, 2, Left}: Zones.GK,
	{-1, 0, Right}: Zones.GK,
	{-1, 1, Right}: Zones.GK,
	{-1, 2, Right}: Zones.GK,
	{0, 0, Right}: Zones.FR,
	{0, 1, Right}: Zones.FC,
	{0, 2, Right}: Zones.FL,
	{1, 0, Right}: Zones.MR,
	{1, 1, Right}: Zones.MC,
	{1, 2, Right}: Zones.ML,
	{2, 0, Right}: Zones.DR,
	{2, 1, Right}: Zones.DC,
	{2, 2, Right}: Zones.DL,

}
func (this *Ball) KickOff(side Side) {
	this.x = 1
	this.y = 1
	this.possessor = side
}
func (this *Ball) Forward() Zone {
	this.y = Dice(3).Throw()
	if this.possessor == Left {
		this.x++
	} else {
		this.x--
	}
	return this.Zone()
}
func (this *Ball) Zone() Zone {
	if this.possessor == Left {
		return field[*this]
	} else {
		return field[*this]
	}
}


