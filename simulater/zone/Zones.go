package zone

import (
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/simulater/zone/effect"
)

type Zones map[name.Name]*Zone


func ZonesFrom(entries Entries) Zones {
	zones := Zones{}
	for _, entry := range entries {
		for _, zoneName := range effect.ConstPositionZoneEffect.ZoneNamesEffectedBy(entry.Position.Name) {
			zones.Get(zoneName).Add(entry)
		}
	}
	return zones
}


func (this Zones) Set(point int) {
	for _, zone := range this {
		for _, entry := range zone.Entries {
			entry.Player.Attributes.Set(point)
		}
	}
}

func (this Zones) Get(name name.Name) *Zone {
	zone, ok := this[name]
	if ok == false {
		zone = &Zone{
			Name:  name,
			Entries: Entries{},
		}
		this[name] = zone
	}
	return zone
}

func (this Zones) Clone() Zones {
	clone := Zones{}
	for key, value := range this {
		clone[key] = value.Clone()
	}
	return clone
}

