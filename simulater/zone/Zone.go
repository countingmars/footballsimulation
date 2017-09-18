package zone

import (
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base/name"
)

type Zone struct {
	Name  name.Name
	Entries Entries

	ability           stats.Ability
}
func (this *Zone) Add(role *Entry) {
	this.Entries = append(this.Entries, role)
}
func (this *Zone) Ability() stats.Ability {
	if len(this.ability) > 0 {
		return this.ability
	}
	this.ability = stats.Ability{}
	for _, entry := range this.Entries {
		this.ability.Add(entry.Player.Ability)
	}
	return this.ability
}

func (this *Zone) calculateAttr(key name.Name, factor float32) float32 {
	attribute, ok := this.Ability()[key]
	if ok {
		return float32(attribute.Point) * factor
	} else {
		return 0
	}
}
func (this *Zone) Scoring() int {
	var sum float32
	for key, factor := range ScoringFactor.AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	return int(sum)
}
func (this *Zone) Possession() int {
	return this.calculate(PossessionFactors)
}

func (this *Zone) Defence() int {
	return this.calculate(DefenceFactors)
}
func (this *Zone) Offence() int {
	return this.calculate(OffenceFactors)
}

func (this *Zone) calculate(positionAttrFactors ZoneAttrFactors) int {
	var sum float32
	for key, factor := range positionAttrFactors[this.Name].AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	sum *= positionAttrFactors[this.Name].Factor
	return int(sum)
}

func (this *Zone) Sum() int {
	sum := 0
	for _, v := range this.Ability() {
		sum += v.Point
	}
	return sum
}
func (this *Zone) Clone() *Zone {
	return &Zone{
		Name:  this.Name,
		Entries: this.Entries.Clone(),
	}
}

