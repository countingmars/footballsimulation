package effect

import (
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/base/player"
)

type Formula struct {
	Formulas   []AttrFormula
	ZoneFactor float32 `json:"zone_factor"`
}

type PlayerFormula struct {
	Zone name.Name
	Position name.Name
	Player *player.Player
}

func (this PlayerFormula) Offence() float32 {
	return this.solve(OffenceZoneAbilityEffect)
}
func (this PlayerFormula) Defence() float32 {
	return this.solve(DefenceZoneAbilityEffect)
}
func (this PlayerFormula) solve(zoneAbilityEffect ZoneAbilityEffect) float32 {
	var sum float32 = 0
	for attr, attrFactor := range zoneAbilityEffect[this.Zone] {
		sum += AttrFormula{this.Player.Ability[attr], attrFactor}.Solve()
	}
	sum *= PositionZoneEffects[this.Position][this.Zone]
	return sum
}
func (this *Formula) Solve() float32 {
	var sum float32 = 0
	for _, each := range this.Formulas {
		sum += each.Solve()
	}
	sum *= this.ZoneFactor
	return sum
}

func (this *Formula) Add(af AttrFormula) {
	this.Formulas = append(this.Formulas, af)
}


type AttrFormula struct {
	A *stats.Attribute `json:"attribute"`
	F float32 `json:"attribute_factor"`
}

func (this AttrFormula) Solve() float32 {
	return float32(this.A.Point) * this.F
}
