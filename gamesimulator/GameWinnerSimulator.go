package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
	. "github.com/countingmars/fb/math"
)
type GameWinnerSimulator struct {

}
func (this GameWinnerSimulator) Simulate(sim *GameSimulation) {
	homeWins, awayWins := 0, 0
	for _, winPoint := range Dice(sim.Game.Sum()).Throws(100) {
		if sim.Game.HomeAbility() < winPoint {
			awayWins++
		} else {
			homeWins++
		}
	}
	if Math.Iabs(homeWins - awayWins) < 3 {
		sim.Draw = true
	} else if homeWins > awayWins {
		sim.Winner = sim.Game.Home
	} else {
		sim.Winner = sim.Game.Away
	}
}