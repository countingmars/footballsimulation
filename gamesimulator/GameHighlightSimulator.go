package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)

type GameHighlightsSimulator struct {

}

func (this *GameHighlightsSimulator) Simulate(sim *GameSimulation) {
	home := sim.Game.Home.Ability
	away := sim.Game.Home.Ability

	decidePossessor(home, away)
	var highlights []Highlight
	for i := 0; i < sim.Score.Sum(); i++ {
		hl := decideHighlight()
		highlights = append(highlights, hl)
	}
	sim.Highlights = highlights
}
func decidePossessor(home TeamAbility, away TeamAbility) int {
	return home.Possession() + away.Possession()
}
func decideHighlight() Highlight {
	index := Dice(cap(Highlights)).Throw()
	return Highlights[index]
}
