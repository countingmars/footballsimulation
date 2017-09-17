package zone

import (
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/base"
)

type Zones map[base.Name]*Zone

func (this Zones) Set(point int) {
	for _, zone := range this {
		zone.Ability().Set(point)
	}
}

func (this Zones) Get(name base.Name) *Zone {
	zone, ok := this[name]
	if ok == false {
		zone = &Zone{
			Name:              name,
			PositionedPlayers: formation.PositionedPlayers{},
		}
		this[name] = zone
	}
	return zone
}
func (this Zones) Sum() int {
	sum := 0
	for _, zone := range this {
		sum += zone.Sum()
	}
	return sum
}

func (this Zones) Clone() Zones {
	clone := Zones{}
	for key, value := range this {
		clone[key] = value.Clone()
	}
	return clone
}

