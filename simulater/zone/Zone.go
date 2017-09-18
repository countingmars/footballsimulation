package zone

import (
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/simulater/zone/effect"
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
	f := effect.Formula{ScoringFactor.AttrFormulas(this), 1.0}
	return int(f.Solve())
}
func (this *Zone) Possession() int {
	return int(PossessionFactors.Formula(this).Solve())
}

func (this *Zone) Defence() int {

	return int(DefenceFactors.Formula(this).Solve())
}

func (this *Zone) Attribute(name name.Name) *stats.Attribute {
	attribute, ok := this.Ability()[name]
	if ok {
		return attribute
	} else {
		return &stats.Attribute{}
	}
}

func (this *Zone) Sum() int {
	sum := 0
	for _, v := range this.Entries {
		sum += v.Player.Ability.Sum()
	}
	return sum
}
func (this *Zone) Clone() *Zone {
	return &Zone{
		Name:  this.Name,
		Entries: this.Entries.Clone(),
	}
}

