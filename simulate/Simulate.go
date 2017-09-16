package simulate

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/simulate/dice"
)

func Simulate(home *base.Team, away *base.Team) *Simulation {
	first := simulateHalf(home.Ability, away.Ability)
	second := simulateHalf(away.Ability, home.Ability)

	simulation := &Simulation{}
	simulation.Home = home
	simulation.Away = away
	simulation.First = first
	simulation.Second = second
	return simulation
}

func simulateHalf(left base.TeamAbility, right base.TeamAbility) Highlights {
	situation := &Situation{
		Timer: &Timer{},
		Ball:  &Ball{},
		Left:  left,
		Right: right,
	}
	situation.Ball.KickOff()

	var halfSimulation = Highlights{}
	for false == situation.Timer.Over() {
		hl := simulateAttack(situation)
		halfSimulation = append(halfSimulation, hl)

		situation.Side = situation.Side.Reverse()
		situation.Timer.Go()
	}
	return halfSimulation
}
func simulateAttack(situation *Situation) Highlight {
	hl := simulateHighlight(situation)
	hl.Side = situation.Side
	hl.Time = situation.Timer.Now()
	return hl
}

func simulateHighlight(situation *Situation) Highlight {
	for {
		if false == simulateBuildup(situation) {
			return Highlight{}
		}

		situation.Ball.Forward()

		if situation.Ball.CanFinish() {
			finish := simulateFinish(situation)
			if finish.Scored() {
				return simulateGoal()
			} else {
				return simulateShoot()
			}
		}
	}
}
func simulateBuildup(context *Situation) bool {
	offence := context.Offender().Offence(context.Ball.Zone())
	defence := context.Defender().Defence(context.Ball.Zone())

	return dice.Judge(offence, defence)
}
func simulateFinish(situation *Situation) Finish {
	scoring := situation.Offender().Scoring(situation.Ball.Zone())
	point := dice.Throw(scoring)

	if point > scoring / 10 {
		return Finish(point)
	} else {
		return Finish(0)
	}
}
func simulateGoal() Highlight {
	return pick(GoalHighlights)
}
func simulateShoot() Highlight {
	return pick(ShootHighlights)
}
func pick(highlights []Highlight) Highlight {
	index := dice.Throw(cap(highlights))
	return highlights[index]
}