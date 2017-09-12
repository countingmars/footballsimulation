package foundation

import (
)

const (
	GK = "GK"
	DL = "DL"
	DC = "DC"
	DR = "DR"
	ML = "ML"
	MC = "MC"
	MR = "MR"
	FL = "FL"
	FC = "FC"
	FR = "FR"
)

type TeamAbility map[string]ZoneAbility


func (this TeamAbility) effects(effects map[string]ZoneAbility) {
	for name, effect := range effects {
		this[name].Effect(effect)
	}
}
func (this TeamAbility) Clone() TeamAbility {
	clone := TeamAbility{}
	for key, value := range this {
		clone[key] = value.Clone()
	}
	return clone
}
func (this TeamAbility) Inspect() *TeamAbilityInspection {
	lineAbility := []*Attribute{}
	lineAbility = append(lineAbility, &Attribute{"D", this.avg(DL, DC, DR)})
	lineAbility = append(lineAbility, &Attribute{"M", this.avg(ML, MC, MR)})
	lineAbility = append(lineAbility, &Attribute{"F", this.avg(FL, FC, FR)})

	sortable(lineAbility).Sort()

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
func (this TeamAbility) avg(zoneNames ... string) int {
	sum := 0
	for _, name := range zoneNames {
		sum += this[name].Sum()
	}
	return sum / len(zoneNames)
}

func (this TeamAbility) Sum() int {
	sum := 0
	for _, v := range this {
		sum += v.Sum()
	}
	return sum
}

func (this TeamAbility) Best() ZoneAbility {
	var best ZoneAbility
	for _, v := range this {
		if best.Sum() < v.Sum() {
			best = v
		}
	}
	return best
}

func (this TeamAbility) Possession() int {
	sum := 0
	for _, value := range this {
		sum += value.Posession()
	}
	return sum
}

