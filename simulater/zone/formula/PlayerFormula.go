package formula

import (
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/base/player"
	"github.com/countingmars/fb/simulater/zone/effect"
)

type PlayerFormula struct {
	Zone name.Name
	Position name.Name
	Player *player.Player
}

func (this PlayerFormula) Offence() float32 {
	return this.Solve(effect.ZoneAttributeEffectOffence)
}
func (this PlayerFormula) Defence() float32 {
	return this.Solve(effect.ZoneAttributeEffectDefence)
}
func (this PlayerFormula) Solve(zoneAttributePremium effect.ZoneAttributeEffect) float32 {
	var sum float32 = 0
	for attr, factor := range zoneAttributePremium[this.Zone] {
		sum += AttributeFormula{this.Player.Attributes[attr], factor}.Solve()
	}
	sum *= effect.ConstPositionZoneEffect[this.Position][this.Zone]
	return sum
}
