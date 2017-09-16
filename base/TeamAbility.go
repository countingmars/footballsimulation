package base


type TeamAbility map[Zone]ZoneAbility

func (this TeamAbility) Scoring(zone Zone) int {
	return this[zone].Scoring()
}

func (this TeamAbility) Defence(zone Zone) int {
	return this[zone].Defence()
}

func (this TeamAbility) Offence(zone Zone) int {
	return this[zone].Offence()
}


func (this TeamAbility) effects(effects map[Zone]ZoneAbility) {
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
	lineAbility = append(lineAbility, &Attribute{"D", this.avg(Zones.DL, Zones.DC, Zones.DR)})
	lineAbility = append(lineAbility, &Attribute{"M", this.avg(Zones.ML, Zones.MC, Zones.MR)})
	lineAbility = append(lineAbility, &Attribute{"F", this.avg(Zones.FL, Zones.FC, Zones.FR)})

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
func (this TeamAbility) avg(zones ... Zone) int {
	sum := 0
	for _, zone := range zones {
		sum += this[zone].Sum()
	}
	return sum / len(zones)
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

