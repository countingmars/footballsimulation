package gamesimulator

import (
	. "github.com/countingmars/fb/foundation"
)

type GameSimulation struct {
	Game   *Game
	Winner *Team
	Draw   bool
	Score  GameScore
	First  HalfSimulation
	Second HalfSimulation
}

