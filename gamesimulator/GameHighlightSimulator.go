package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)

type GameHighlightsSimulator struct {

}

func (this *GameHighlightsSimulator) Simulate(sim *GameSimulation) {
	var highlights []Highlight
	for i := 0; i < sim.Score.Sum(); i++ {
		hl := decideHighlight()
		highlights = append(highlights, hl)
	}
	sim.Highlights = highlights
}
func decideHighlight() Highlight {
	index := Dice(cap(Highlights)).Throw()
	return Highlights[index]
}
