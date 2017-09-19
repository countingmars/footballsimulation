package zone

import (
	"github.com/countingmars/fb/simulater/zone/formula"
	"github.com/countingmars/fb/simulater/zone/effect"
	"github.com/countingmars/fb/simulater/dice"
)

type ZoneFormula struct {
	Zone *Zone
}
func MakeZoneFormula(zone *Zone) *ZoneFormula {
	return &ZoneFormula{zone}
}
func (this *ZoneFormula) Scoring() int {
	return this.Solve(effect.ZoneAttributeEffectScoring)
}
func (this *ZoneFormula) Offence() int {
	return this.Solve(effect.ZoneAttributeEffectOffence)
}
func (this *ZoneFormula) Defence() int {
	return this.Solve(effect.ZoneAttributeEffectDefence)
}
func (this *ZoneFormula) Solve(zoneAttributeEffect effect.ZoneAttributeEffect) int {
	var sum = 0
	for _, entry := range this.Zone.Entries {
		playerFormula := formula.PlayerFormula{this.Zone.Name, entry.Position.Name, entry.Player}
		max := playerFormula.Solve(zoneAttributeEffect)
		actual := dice.Throw(int(max))
		entry.Stats.LastPoint = actual
		sum += actual
	}
	return sum
}