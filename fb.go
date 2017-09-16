package main

import (
	"fmt"
	. "github.com/countingmars/fb/foundation"
	. "github.com/countingmars/fb/simulator"
	. "github.com/countingmars/fb/commentor"
	. "github.com/countingmars/fb/data"
)

func main() {
	fmt.Println("fb started")

	model := LoadTeam("./data/json/team-perfect.json")
	ManchesterUnited := model.Clone()
	ManchesterUnited.Name = "Manchester United"
	RealMadrid := model.Clone()
	RealMadrid.Name = "Real Madrid"

	game := &Game {ManchesterUnited, RealMadrid}
	simulator := GameSimulator{}

	manuWinCount := 0
	realWinCount := 0
	drawCount := 0
	for i := 1; i <= 100; i++ {
		simulation := simulator.Simulate(game)
		if simulation.Draw == true {
			drawCount++
		} else if simulation.Winner.Equals(ManchesterUnited) {
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
