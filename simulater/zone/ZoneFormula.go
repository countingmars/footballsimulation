package zone

import (
	"github.com/countingmars/fb/simulater/zone/formula"
	"github.com/countingmars/fb/simulater/zone/effect"
)

func MakeZoneFormula(zone *Zone) *ZoneFormula {
	zf := &ZoneFormula{}
	zf.Zone(zone)
	return zf
}

type ZoneFormula struct {
	playerFormulas []formula.PlayerFormula
}


func (this *ZoneFormula) Zone(zone *Zone) {
	for _, entry := range zone.Entries {
		playerFormula := formula.PlayerFormula{zone.Name, entry.Position.Name, entry.Player}
		this.playerFormulas = append(this.playerFormulas, playerFormula)
	}
}
func (this *ZoneFormula) Scoring() float32 {
	return this.Solve(effect.ZoneAttributeEffectScoring)
}
func (this *ZoneFormula) Offence() float32 {
	return this.Solve(effect.ZoneAttributeEffectOffence)
}
func (this *ZoneFormula) Defence() float32 {
	return this.Solve(effect.ZoneAttributeEffectDefence)
}
func (this *ZoneFormula) Solve(zoneAttributePremium effect.ZoneAttributeEffect) float32 {
	var sum float32
	for _, each := range this.playerFormulas {
		sum += each.Solve(zoneAttributePremium)
	}
	return sum
}