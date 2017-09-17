package stats

import "github.com/countingmars/fb/base"

type Ability map[base.Name]*Attribute

func (this Ability) Add(another Ability) {
	for key, value := range another {
		attribute, ok := this[key]
		if !ok {
			attribute = &Attribute{Name: key}
			this[key] = attribute
		}
		this[key].Point += value.Point
	}
}

func (this Ability) Clone() Ability {
	clone := Ability{}
	for name, attr := range this {
		attrClone := attr.Clone()
		clone[name] = &attrClone
	}
	return clone
}

func (this Ability) Set(point int) {
	for _, attr := range this {
		attr.Point = point
	}
}
func (this Ability) Sum() int {
	sum := 0
	for _, v := range this {
		sum += v.Point
	}
	return sum
}

