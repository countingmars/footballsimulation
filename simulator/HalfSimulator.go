package simulator

import (
	. "github.com/countingmars/fb/foundation"
)



type HalfSimulator struct {
	Timer *Timer
	Ball *Ball
	Side Side
	Left TeamAbility
	Right TeamAbility
}
func NewHalfSimulator(left TeamAbility, right TeamAbility) HalfSimulator {
	return HalfSimulator{
		Timer: &Timer{},
		Ball: &Ball{},
		Side: Left,
		Left: left,
		Right: right,
	}
}
func (this *HalfSimulator) Simulate() HalfSimulation {
	this.Ball.KickOff(Left)

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
	hl := this.attack(this.buildAttackContext())
	hl.Side = this.Side
	hl.Time = this.Timer.Now()
	return hl
}
func (this *HalfSimulator) buildAttackContext() *AttackSituation {
	if this.Side == Left {
		return &AttackSituation{this.Left, this.Right, this.Ball, false, false, HighlightSimulator{},}
	} else {
		return &AttackSituation{this.Right, this.Left, this.Ball, false, false, HighlightSimulator{},}
	}
}
func (this *HalfSimulator) attack(context *AttackSituation) Highlight {
	return context.Play()

}
