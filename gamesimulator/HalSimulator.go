package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
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
		Ball: &Ball{Position: MC, Chance: false },
		Side: Left,
		Left: left,
		Right: right,
		HighlightSimulator: HighlightSimulator{},
	}
}
func (this *HalfSimulator) Simulate() HalfSimulation {
	var halfSimulation = HalfSimulation{}
	for false == this.Timer.Over() {
		hl := this.simulateAttack()
		halfSimulation = append(halfSimulation, hl)

		this.Side = this.Side.Reverse()
		this.Timer.Go()
	}
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
		buildup := this.simulateBuildup(offender, defender)
		if buildup.Ball.Chance {
			penetration := this.simulatePenetration(offender, defender)
			if penetration.Scored() {
				return this.HighlightSimulator.SimulateGoal()
			} else {
				return this.HighlightSimulator.SimulateShoot()
			}
		}
		if &buildup.Possessor == &defender {
			return Highlight{}
		}
	}
}

func (this *HalfSimulator) simulateBuildup(offender TeamAbility, defender TeamAbility) Buildup {
	sum := offender[this.Ball.Position].Offence() + defender[this.Ball.Position].Defence()
	point := Dice(sum).Throw()
	if point < offender[this.Ball.Position].Offence() {
		this.Ball.Forward()
		return Buildup{ offender, this.Ball }
	} else {
		return Buildup{ defender, this.Ball }
	}
}

func (this *HalfSimulator)  simulatePenetration(offender TeamAbility, defender TeamAbility) Penetration {
	scoring := offender[this.Ball.Position].Scoring()
	point := Dice(scoring).Throw()

	if point > scoring / 10 {
		return Penetration{ point }
	} else {
		return Penetration{}
	}
}
