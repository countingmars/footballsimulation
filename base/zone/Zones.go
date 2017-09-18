package zone

import (
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/base/name"
)

type Zones map[name.Name]*Zone


func ZonesFrom(aFormation formation.Formation) Zones {
	zones := Zones{}
	for _, each := range aFormation {
		zoneNames := resolveZoneNames(each.Position)
		for _, zoneName := range zoneNames {
			zones.Get(zoneName).Add(each)
		}
	}
	return zones
}


func (this Zones) Set(point int) {
	for _, zone := range this {
		zone.Ability().Set(point)
	}
}

func (this Zones) Get(name name.Name) *Zone {
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

