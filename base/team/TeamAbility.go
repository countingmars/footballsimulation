package team

import (
	"github.com/countingmars/fb/base/attr"
	"github.com/countingmars/fb/base/zone"
)

type Ability map[zone.Name]zone.Ability

func (this Ability) Set(point int) {
	for _, each := range this {
		each.Set(point)
	}
}


func (this Ability) Scoring(zone zone.Name) int {
	return this[zone].Scoring()
}

func (this Ability) Defence(zone zone.Name) int {
	return this[zone].Defence()
}

func (this Ability) Offence(zone zone.Name) int {
	return this[zone].Offence()
}


func (this Ability) effects(effects map[zone.Name]zone.Ability) {
	for name, effect := range effects {
		this[name].Effect(effect)
	}
}
func (this Ability) Clone() Ability {
	clone := Ability{}
	for key, value := range this {
		clone[key] = value.Clone()
	}
	return clone
}
func (this Ability) Inspect() *TeamAbilityInspection {
	lineAbility := []*attr.Attribute{}
	lineAbility = append(lineAbility, &attr.Attribute{"D", this.avg(zone.Names.DL, zone.Names.DC, zone.Names.DR)})
	lineAbility = append(lineAbility, &attr.Attribute{"M", this.avg(zone.Names.ML, zone.Names.MC, zone.Names.MR)})
	lineAbility = append(lineAbility, &attr.Attribute{"F", this.avg(zone.Names.FL, zone.Names.FC, zone.Names.FR)})

	attr.Attributes(lineAbility).Sort()

	inspection := &TeamAbilityInspection{
		Strength: lineAbility[len(lineAbility) - 1],
		Weakness: lineAbility[0],
	}

	switch inspection.Strength.Name {
		case "D": inspection.Style = TAI_DEFENCE
		case "M": inspection.Style = TAI_PASS
		case "F": inspection.Style = TAI_ATTACK
	}
	if inspection.Diff() <= 2 {
		inspection.Style = TAI_BALANCE
	}

	return inspection
}
func (this Ability) avg(zones ... zone.Name) int {
	sum := 0
	for _, zone := range zones {
		sum += this[zone].Sum()
	}
	return sum / len(zones)
}

func (this Ability) Sum() int {
	sum := 0
	for _, v := range this {
		sum += v.Sum()
	}
	return sum
}

func (this Ability) Best() zone.Ability {
	var best zone.Ability
	for _, v := range this {
		if best.Sum() < v.Sum() {
			best = v
		}
	}
	return best
}

func (this Ability) Possession() int {
	sum := 0
	for _, value := range this {
		sum += value.Posession()
	}
	return sum
}

