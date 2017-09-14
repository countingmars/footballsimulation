package gamesimulator

import (
	"strings"
	. "github.com/countingmars/fb/foundation"
)

type Ball struct {
	Position string
	Chance bool
}
func (this *Ball) Forward() {
	this.Chance = false

	point := Dice(3).Throw()
	lrc := []string{ "L", "R", "C" }
	if 0 == strings.Index(this.Position, "D") {
		this.Position = "M" + lrc[point]
	} else if 0 == strings.Index(this.Position, "M") {
		this.Position = "F" + lrc[point]
	} else { // when offenders get a chance to shoot
		this.Chance = true
	}
}


