package stats

import "github.com/countingmars/fb/base/name"

type Attributes map[name.Name]*Attribute

func (this Attributes) Add(another Attributes) {
	for key, value := range another {
		attribute, ok := this[key]
		if !ok {
			attribute = &Attribute{Name: key}
			this[key] = attribute
		}
		this[key].Point += value.Point
	}
}

func (this Attributes) Clone() Attributes {
	clone := Attributes{}
	for name, attr := range this {
		attrClone := attr.Clone()
		clone[name] = &attrClone
	}
	return clone
}

func (this Attributes) Set(point int) {
	for _, attr := range this {
		attr.Point = point
	}
}
func (this Attributes) Sum() int {
	sum := 0
	for _, v := range this {
		sum += v.Point
	}
	return sum
}

