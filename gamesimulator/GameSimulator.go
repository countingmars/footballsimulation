package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)
type GameSimulator struct {
	highlight *GameHighlightsSimulator
}
func CreateGameSimulator() *GameSimulator {
	return &GameSimulator{
		highlight: &GameHighlightsSimulator{},
	}
}
func (this *GameSimulator) Simulate(game *Game) *GameSimulation {
	simulation := &GameSimulation {
		Game: game,
	}

	this.highlight.Simulate(simulation)

	return simulation	
}

