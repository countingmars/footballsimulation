package zone

import (
	"sort"
	"github.com/countingmars/fb/base/attr"
)

type Ability struct {
	Name       Name
	Attributes map[attr.Name]*attr.Attribute
}

func (this Ability) Avg() int {
	return this.Sum() / len(this.Attributes)
}

func (this Ability) Set(point int) {
	for _, each := range this.Attributes {
		each.Point = point
	}
}
func (this Ability) calculateAttr(key attr.Name, factor float32) float32 {
	attribute, ok := this.Attributes[key]
	if ok {
		return float32(attribute.Point) * factor
	} else {
		return 0
	}
}
func (this Ability) Scoring() int {
	var sum float32
	for key, factor := range ScoringFactor.AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	return int(sum)
}
func (this Ability) Posession() int {
	return this.calculate(PossessionFactors)
}
func (this Ability) Defence() int {
	return this.calculate(DefenceFactors)
}
func (this Ability) Offence() int {
	return this.calculate(OffenceFactors)
}
func (this Ability) calculate(positionAttrFactors ZoneAttrFactors) int {
	var sum float32
	for key, factor := range positionAttrFactors[this.Name].AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	sum *= positionAttrFactors[this.Name].Factor
	return int(sum)
}

func (this Ability) Sum() int {
	sum := 0
	for _, v := range this.Attributes {
		sum += v.Point
	}
	return sum
}

func (this Ability) Sort() []*attr.Attribute {
	array := []*attr.Attribute{}
	for _, el := range this.Attributes {
		array = append(array, el)
	}
	sort.Sort(attr.Attributes(array))
	return array
}
func (this Ability) Effect(debuff Ability) {

}
func (this Ability) Clone() Ability {
	clone := Ability{ Name: this.Name, Attributes: map[attr.Name]*attr.Attribute{} }
	for k, v := range this.Attributes {
		ability := *v
		clone.Attributes[k] = &ability
	}
	return clone
}
func (this Ability) StrongerThan(other Ability) bool {
	return this.Sum() > other.Sum()
}
