package main

import (
	"fmt"
	. "github.com/countingmars/fb/foundation"
	. "github.com/countingmars/fb/gamesimulator"
	. "github.com/countingmars/fb/commentor"
	. "github.com/countingmars/fb/data"
)

func main() {
	fmt.Println("fb started")

	manchester := LoadTeam("./data/json/team-manu.json")
	madrid := LoadTeam("./data/json/team-real.json")

	game := &Game {manchester, madrid}
	simulator := CreateGameSimulator()

	manuWinCount := 0
	realWinCount := 0
	drawCount := 0
	for i := 1; i <= 100; i++ {
		simulation := simulator.Simulate(game)
		if simulation.Draw == true {
			drawCount++
		} else if simulation.Winner.Equals(manchester) {
			manuWinCount++ 
		} else { 
			realWinCount++
		}
		commentor := Commentor{}
		fmt.Println(commentor.Comment(simulation).Message)
	}
	fmt.Println(manuWinCount, " vs ", realWinCount, " draw : ", drawCount)
	fmt.Println("fb completed")
}
