package main

import (
	"fmt"
	"github.com/countingmars/fb/simulate"
	"github.com/countingmars/fb/commentor"
	"github.com/countingmars/fb/data"
)

func main() {
	fmt.Println("fb started")

	model := data.LoadTeam("./data/json/team-perfect.json")
	ManchesterUnited := model.Clone()
	ManchesterUnited.Name = "Manchester United"
	RealMadrid := model.Clone()
	RealMadrid.Name = "Real Madrid"

	manuWinCount := 0
	realWinCount := 0
	drawCount := 0
	for i := 1; i <= 100; i++ {
		simulation := simulate.Simulate(ManchesterUnited, RealMadrid)
		if simulation.Draw() {
			drawCount++
		} else if simulation.Win() {
			manuWinCount++ 
		} else { 
			realWinCount++
		}
		fmt.Println(commentor.Comment(simulation))
	}
	fmt.Println(manuWinCount, " vs ", realWinCount, " draw : ", drawCount)
	fmt.Println("fb completed")
}
