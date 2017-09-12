package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)
type GameSimulator struct {
	winner *GameWinnerSimulator
	goal *GameGoalsSimulator
	highlight *GameHighlightsSimulator
}
func CreateGameSimulator() *GameSimulator {
	return &GameSimulator{
		winner: &GameWinnerSimulator{},
		goal: &GameGoalsSimulator{},
		highlight: &GameHighlightsSimulator{},
	}
}
func (this *GameSimulator) Simulate(game *Game) *GameSimulation {
	simulation := &GameSimulation {
		Game: game,
	}

	this.winner.Simulate(simulation)
	this.goal.Simulate(simulation)
	this.highlight.Simulate(simulation)

	return simulation	
}

