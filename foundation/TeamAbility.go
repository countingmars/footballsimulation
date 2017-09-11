package foundation

import (
)

const (
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

type Effect struct {
	Bad ZoneAbility
	Good ZoneAbility
	Factor float32
}
func (this Effect) Calculate() int {
	diff := this.Good.Sum() - this.Bad.Sum()
	if diff > 0 {
		factor := float32(diff) * this.Factor
		if factor < 1 {
			return 1
		} else {
			return int(factor)
		}
	} else {
		return 0
	}
}

func (this TeamAbility) calculateEffectsBy(other TeamAbility) map[string]int {
	debuff := map[string]int{}
	debuff[DL] =+ Effect{this[DL], other[FR], 0.35}.Calculate()
	debuff[DL] =+ Effect{this[DL], other[MR], 0.15}.Calculate()
	debuff[DL] =+ Effect{this[DL], other[MC], 0.10}.Calculate()
	debuff[DL] =+ Effect{this[DL], other[FC], 0.15}.Calculate()

	debuff[DC] =+ Effect{this[DC], other[FC], 0.35}.Calculate()
	debuff[DC] =+ Effect{this[DC], other[MC], 0.25}.Calculate()
	debuff[DC] =+ Effect{this[DC], other[FR], 0.20}.Calculate()
	debuff[DC] =+ Effect{this[DC], other[FL], 0.20}.Calculate()

	debuff[DR] =+ Effect{this[DR], other[FL], 0.35}.Calculate()
	debuff[DR] =+ Effect{this[DR], other[ML], 0.15}.Calculate()
	debuff[DR] =+ Effect{this[DR], other[MC], 0.10}.Calculate()
	debuff[DR] =+ Effect{this[DR], other[FC], 0.15}.Calculate()

	debuff[ML] =+ Effect{this[ML], other[MR], 0.35}.Calculate()
	debuff[ML] =+ Effect{this[ML], other[DR], 0.25}.Calculate()
	debuff[ML] =+ Effect{this[ML], other[FR], 0.10}.Calculate()
	debuff[ML] =+ Effect{this[ML], other[MC], 0.15}.Calculate()

	debuff[MC] =+ Effect{this[MC], other[MC], 0.35}.Calculate()
	debuff[MC] =+ Effect{this[MC], other[ML], 0.20}.Calculate()
	debuff[MC] =+ Effect{this[MC], other[MR], 0.20}.Calculate()
	debuff[MC] =+ Effect{this[MC], other[DC], 0.05}.Calculate()
	debuff[MC] =+ Effect{this[MC], other[FC], 0.05}.Calculate()

	debuff[MR] =+ Effect{this[MR], other[ML], 0.35}.Calculate()
	debuff[MR] =+ Effect{this[MR], other[DL], 0.25}.Calculate()
	debuff[MR] =+ Effect{this[MR], other[FL], 0.10}.Calculate()
	debuff[MR] =+ Effect{this[MR], other[MC], 0.15}.Calculate()

	debuff[FL] =+ Effect{this[FL], other[DR], 0.35}.Calculate()
	debuff[FL] =+ Effect{this[FL], other[MR], 0.25}.Calculate()
	debuff[FL] =+ Effect{this[FL], other[MC], 0.05}.Calculate()
	debuff[FL] =+ Effect{this[FL], other[DC], 0.20}.Calculate()

	debuff[FC] =+ Effect{this[FC], other[DC], 0.35}.Calculate()
	debuff[FC] =+ Effect{this[FC], other[MC], 0.05}.Calculate()
	debuff[FC] =+ Effect{this[FC], other[DL], 0.15}.Calculate()
	debuff[FC] =+ Effect{this[FC], other[DR], 0.15}.Calculate()

	debuff[FR] =+ Effect{this[FR], other[DL], 0.35}.Calculate()
	debuff[FR] =+ Effect{this[FR], other[ML], 0.25}.Calculate()
	debuff[FR] =+ Effect{this[FR], other[MC], 0.05}.Calculate()
	debuff[FR] =+ Effect{this[FR], other[DC], 0.20}.Calculate()

	return debuff
}
func (this TeamAbility) calculateEffectsBySelf() map[string]int {
	negative := map[string]int{}

	// the effect of defenders is a lot
	negative[DC] += Effect{this[DL], this[DC], 0.45}.Calculate()
	negative[ML] += Effect{this[DL], this[ML], 0.35}.Calculate()
	negative[MC] += Effect{this[DL], this[MC], 0.25}.Calculate()

	negative[DL] += Effect{this[DC], this[DL], 0.35}.Calculate()
	negative[DR] += Effect{this[DC], this[DR], 0.35}.Calculate()
	negative[MC] += Effect{this[DC], this[MC], 0.45}.Calculate()

	negative[DC] += Effect{this[DR], this[DC], 0.45}.Calculate()
	negative[MR] += Effect{this[DR], this[MR], 0.35}.Calculate()
	negative[MC] += Effect{this[DR], this[MC], 0.25}.Calculate()

	negative[MC] += Effect{this[ML], this[MC], 0.35}.Calculate()
	negative[DL] += Effect{this[ML], this[DL], 0.15}.Calculate()
	negative[FL] += Effect{this[ML], this[FL], 0.25}.Calculate()

	// the effect of center-mid is wide spread
	negative[ML] += Effect{this[MC], this[ML], 0.20}.Calculate()
	negative[MR] += Effect{this[MC], this[MR], 0.20}.Calculate()
	negative[DC] += Effect{this[MC], this[DC], 0.20}.Calculate()
	negative[FC] += Effect{this[MC], this[FC], 0.35}.Calculate()
	negative[FL] += Effect{this[MC], this[FL], 0.20}.Calculate()
	negative[FR] += Effect{this[MC], this[FR], 0.20}.Calculate()

	negative[MC] += Effect{this[MR], this[MC], 0.35}.Calculate()
	negative[DR] += Effect{this[MR], this[DR], 0.15}.Calculate()
	negative[FR] += Effect{this[MR], this[FR], 0.25}.Calculate()

	// the effect of forward players is small,
	// but the creative and decision are going to be important,
	// which make opponents confusing

	negative[FC] += Effect{this[FL], this[FC], 0.25}.Calculate()
	negative[ML] += Effect{this[FL], this[ML], 0.15}.Calculate()

	negative[FL] += Effect{this[FC], this[FL], 0.25}.Calculate()
	negative[MC] += Effect{this[FC], this[MC], 0.15}.Calculate()
	negative[FR] += Effect{this[FC], this[FR], 0.15}.Calculate()

	negative[FC] += Effect{this[FR], this[FC], 0.25}.Calculate()
	negative[MR] += Effect{this[FR], this[MR], 0.15}.Calculate()

	return negative
}
func (this TeamAbility) effects(effects map[string]int) {
	for k, v := range effects {
		this[k].Debuff(v)
	}
}
func (this TeamAbility) Clone() TeamAbility {
	clone := TeamAbility{}
	for k, v := range this {
		clone[k] = v.Clone()
	}
	return clone
}
func (this TeamAbility) AmendBySelf() TeamAbility {
	negative := this.calculateEffectsBySelf()
	effected := this.Clone()

	effected.effects(negative)
	return effected
}
func (this TeamAbility) AmendBy(other TeamAbility) TeamAbility {
	negative := this.calculateEffectsBy(other)
	effected := this.Clone()

	effected.effects(negative)
	return effected
}
func (this TeamAbility) Inspect() *TeamAbilityInspection {
	lineAbility := []*Ability{}
	lineAbility = append(lineAbility, &Ability{"D", this.avg(DL, DC, DR)})
	lineAbility = append(lineAbility, &Ability{"M", this.avg(ML, MC, MR)})
	lineAbility = append(lineAbility, &Ability{"F", this.avg(FL, FC, FR)})

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

