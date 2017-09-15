package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
	"fmt"
)



type HalfSimulator struct {
	Timer *Timer
	Ball *Ball
	Side Side
	Left TeamAbility
	Right TeamAbility
	HighlightSimulator HighlightSimulator
}
func NewHalfSimulator(left TeamAbility, right TeamAbility) HalfSimulator {
	return HalfSimulator{
		Timer: &Timer{},
		Ball: &Ball{},
		Side: Left,
		Left: left,
		Right: right,
		HighlightSimulator: HighlightSimulator{},
	}
}
func (this *HalfSimulator) Simulate() HalfSimulation {
	this.Ball.KickOff()
	var halfSimulation = HalfSimulation{}
	for false == this.Timer.Over() {
		hl := this.simulateAttack()
		halfSimulation = append(halfSimulation, hl)

		this.Side = this.Side.Reverse()
		this.Timer.Go()

	}
	fmt.Println()
	return halfSimulation
}
func (this *HalfSimulator) simulateAttack() Highlight {
	var hl Highlight
	if this.Side == Left {
		hl = this.simulateAttackFromLeft()
	} else {
		hl = this.simulateAttackFromRight()
	}
	hl.Side = this.Side
	hl.Time = this.Timer.Now()
	return hl
}
func (this *HalfSimulator) simulateAttackFromLeft() Highlight {
	return this.simulateAttackFromTo(this.Left, this.Right)
}
func (this *HalfSimulator) simulateAttackFromRight() Highlight {
	return this.simulateAttackFromTo(this.Right, this.Left)
}
func (this *HalfSimulator) simulateAttackFromTo(offender TeamAbility, defender TeamAbility) Highlight {
	for {
		success := SimulateBuildup(offender, defender, this.Ball, this.Side)
		if success {
			if this.Ball.Chance(this.Side) {
				penetration := this.simulatePenetration(offender, defender)
				if penetration.Scored() {
					return this.HighlightSimulator.SimulateGoal()
				} else {
					return this.HighlightSimulator.SimulateShoot()
				}
			}
		} else {
			return Highlight{}
		}
	}
}

func SimulateBuildup(offender TeamAbility, defender TeamAbility, ball *Ball, side Side) bool {
	position := ball.Position(side)

	sum := offender[position].Offence() + defender[position].Defence()
	point := Dice(sum).Throw()

	if point < offender[position].Offence() {
		ball.Forward(side)
		return true
	} else {
		return false
	}
}

func (this *HalfSimulator)  simulatePenetration(offender TeamAbility, defender TeamAbility) Penetration {
	scoring := offender[this.Ball.Position(this.Side)].Scoring()
	point := Dice(scoring).Throw()

	if point > scoring / 10 {
		return Penetration{ point }
	} else {
		return Penetration{}
	}
}
