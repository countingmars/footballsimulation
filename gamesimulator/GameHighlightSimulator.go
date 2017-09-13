package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
	"strings"
)

type GameHighlightsSimulator struct {

}

func (this *GameHighlightsSimulator) Simulate(sim *GameSimulation) {
	firstHalfHighlights := simulateHalf(sim.Game.Home.Ability, sim.Game.Away.Ability)
	secondHighlights := simulateHalf(sim.Game.Away.Ability, sim.Game.Home.Ability)

	sim.Highlights = append(firstHalfHighlights, secondHighlights...)
}
type Buildup struct {
	Possessor TeamAbility
	Ball *Ball
}
type Ball struct {
	Position string
	Chance bool
}
func simulateHalf(offender TeamAbility, defender TeamAbility) []Highlight {
	const HALF = 45
	const FIVE_MINUTES = 5
	const CHANCES = HALF / FIVE_MINUTES

	var ball = &Ball { Position: MC, Chance: false }
	var highlights = []Highlight{}
	for i := 0; i < CHANCES; i++ {
		highlight, ok := simulateAttack(offender, defender, ball)
		if ok {
			highlights = append(highlights, highlight)
		}
		tmp := offender
		offender = defender
		defender = tmp
	}
	return highlights
}
func simulateAttack(offender TeamAbility, defender TeamAbility, ball *Ball) (Highlight, bool) {
	for {
		buildup := simulateBuildup(offender, defender, ball)
		if buildup.Ball.Chance {
			penetration := simulatePenetration(offender, defender, ball)
			if penetration.Scored() {
				index := Dice(cap(GoalScenarios)).Throw()
				return GoalScenarios[index], true
			} else {

				return Highlight{}, false
			}
		}
		if &buildup.Possessor == &defender {
			return Highlight{}, false
		}
	}
}
func simulatePenetration(offender TeamAbility, defender TeamAbility, ball *Ball) Penetration {
	scoring := offender[ball.Position].Scoring()
	point := Dice(scoring).Throw()

	if point > scoring / 10 {
		return Penetration{ point }
	} else {
		return Penetration{}
	}
}
func simulateBuildup(offender TeamAbility, defender TeamAbility, ball *Ball) Buildup {
	sum := offender[ball.Position].Offence() + defender[ball.Position].Defence()
	point := Dice(sum).Throw()
	if point < offender[ball.Position].Offence() {
		ball.Forward()
		return Buildup{ offender, ball }
	} else {
		return Buildup{ defender, ball }
	}
}
func (this *Ball) Forward() {
	this.Chance = false

	point := Dice(3).Throw()
	lrc := []string{ "L", "R", "C" }
	if 0 == strings.Index(this.Position, "D") {
		this.Position = "M" + lrc[point]
	} else if 0 == strings.Index(this.Position, "M") {
		this.Position = "F" + lrc[point]
	} else { // when offenders get a chance to shoot
		this.Chance = true
	}
}

