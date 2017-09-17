package main

import (
	"fmt"
	"github.com/countingmars/fb/simulater"
	"github.com/countingmars/fb/commentor"
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/base/formation"
)

func main() {
	fmt.Println("fb started")

	ManchesterUnited := &team.Team{}
	ManchesterUnited.Name = "Manchester United"
	ManchesterUnited.Formation = formation.Test442()

	RealMadrid := ManchesterUnited.Clone()
	RealMadrid.Name = "Real Madrid"
	RealMadrid.Formation.Set(10)

	manuWinCount := 0
	realWinCount := 0
	drawCount := 0
	for i := 1; i <= 100; i++ {
		simulation := simulater.Simulate(ManchesterUnited, RealMadrid)
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


