package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)
type HighlightSimulator struct {

}


func (this HighlightSimulator) SimulateGoal() Highlight {
	return this.simulate(GoalHighlights)
}
func (this HighlightSimulator) SimulateShoot() Highlight {
	return this.simulate(ShootHighlights)
}
func (this HighlightSimulator) simulate(highlights []Highlight) Highlight {
	index := Dice(cap(highlights)).Throw()
	return highlights[index]
}
