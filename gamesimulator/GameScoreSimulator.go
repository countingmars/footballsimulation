package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)
type GameGoalsSimulator struct {

}

func (this *GameGoalsSimulator) Simulate(sim *GameSimulation) {
	if sim.Draw {
		goals := Dice(10).Throw()
		if goals <= 2 {
			sim.Score = GameScore { 0, 0 }
		} else if goals <= 5 {
			sim.Score = GameScore { 1, 1 }
		} else if goals <= 8 {
			sim.Score = GameScore { 2, 2 }
		} else {
			sim.Score = GameScore { 3, 3 }
		}
	} else {
		var winnerGoals, loserGoals int
		goals := Dice(100).Throw()
		if goals <= 40 {
			winnerGoals = 1
		} else if goals <= 70 {
			winnerGoals = 2
		} else if goals <= 85 {
			winnerGoals = 3
		} else if goals <= 90 {
			winnerGoals = 4
		} else if goals <= 95 {
			winnerGoals = 5
		} else {
			winnerGoals = 5 + 100 - goals
		}
		loserGoals = Dice(winnerGoals).Throw()
		if sim.Winner == sim.Game.Home {
			sim.Score = GameScore { winnerGoals, loserGoals }
		} else {
			sim.Score = GameScore { loserGoals, winnerGoals }
		}
	}
}