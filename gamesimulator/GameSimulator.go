package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)
type GameSimulator struct {
}
func (this *GameSimulator) Simulate(game *Game) *GameSimulation {
	simulation := &GameSimulation {
		Game: game,
	}

	firstHalfSimulator := NewHalfSimulator(game.Home.Ability, game.Away.Ability);
	firstHalf := firstHalfSimulator.Simulate()

	secondHalfSimulator := NewHalfSimulator(game.Away.Ability, game.Home.Ability)
	secondHalf := secondHalfSimulator.Simulate()


	simulation.Score = GameScore{
		Home: firstHalf.Goals(Left) + secondHalf.Goals(Right),
		Away: firstHalf.Goals(Right) + secondHalf.Goals(Left),
	}

	if simulation.Score.Home == simulation.Score.Away {
		simulation.Draw = true
	} else if simulation.Score.Home > simulation.Score.Away {
		simulation.Winner = simulation.Game.Home
	} else {
		simulation.Winner = simulation.Game.Away
	}
	simulation.First = firstHalf
	simulation.Second = secondHalf
	return simulation	
}

