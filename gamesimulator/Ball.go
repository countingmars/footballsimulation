package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)

type Ball struct {
	x int
	y int
}
var leftField = map[Ball]string {
	{0, 0}: "DR",
	{0, 1}: "DC",
	{0, 2}: "DL",
	{1, 0}: "MR",
	{1, 1}: "MC",
	{1, 2}: "ML",
	{2, 0}: "FR",
	{2, 1}: "FC",
	{2, 2}: "FL",
}
var rightField = map[Ball]string {
	{0, 0}: "FR",
	{0, 1}: "FC",
	{0, 2}: "FL",
	{1, 0}: "MR",
	{1, 1}: "MC",
	{1, 2}: "ML",
	{2, 0}: "DR",
	{2, 1}: "DC",
	{2, 2}: "DL",
}
func (this *Ball) KickOff() {
	this.x = 1
	this.y = 1
}
func (this *Ball) Forward(side Side) {
	this.y = Dice(3).Throw()
	if side == Left {
		this.x++
	} else {
		this.x--
	}
}
func (this *Ball) Position(side Side) string {
	if side == Left {
		return leftField[*this]
	} else {
		return rightField[*this]
	}
}
func (this *Ball) Chance(side Side) bool {
	if side == Left {
		return this.x == 3
	} else {
		return this.x == -1
	}
}


