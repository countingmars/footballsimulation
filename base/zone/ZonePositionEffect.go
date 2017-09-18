package zone

import (
	"github.com/countingmars/fb/base/name"
)

type ZoneEffectTable map[name.Name]float32
func (this ZoneEffectTable) Contains(name name.Name) bool {
	_, ok := this[name]
	return ok
}

type PositionEffectTable map[name.Name]ZoneEffectTable
func (this PositionEffectTable) ZoneNamesFor(positionName name.Name) name.Names {
	zoneNames := name.Names{}
	for zoneName, _ := range this[positionName] {
		zoneNames = append(zoneNames, zoneName)
	}
	return zoneNames
}
var PET = PositionEffectTable{
	name.GK: ZoneEffectTable{
		name.GK: 1.0,
	},
	name.DC: ZoneEffectTable{
		name.DC: 1.0,
		name.DL: 0.4,
		name.DR: 0.4,
		name.MC: 0.2,
	},
	name.DL: ZoneEffectTable{
		name.DL: 1.0,
		name.ML: 0.5,
		name.DC: 0.5,
	},
	name.DR: ZoneEffectTable{
		name.DR: 1.0,
		name.MR: 0.5,
		name.DC: 0.5,
	},
	name.DMC: ZoneEffectTable{
		name.MC: 0.5,
		name.DC: 0.5,
		name.DL: 0.5,
		name.DR: 0.5,
	},
	name.WBL: ZoneEffectTable{
		name.DL: 0.5,
		name.ML: 0.5,
		name.FL: 0.5,
		name.MC: 0.5,
	},
	name.WBR: ZoneEffectTable{
		name.DR: 0.5,
		name.MR: 0.5,
		name.FR: 0.5,
		name.MC: 0.5,
	},
	name.MC: ZoneEffectTable{
		name.MC: 1.0,
		name.FC: 0.5,
		name.ML: 0.5,
		name.MR: 0.5,
	},
	name.ML: ZoneEffectTable{
		name.ML: 1.0,
		name.DL: 0.5,
		name.FL: 0.5,
		name.MC: 0.5,
	},
	name.AMC: ZoneEffectTable{
		name.MC: 0.5,
		name.FC: 0.5,
		name.FL: 0.5,
		name.FR: 0.5,
	},
	name.AML: ZoneEffectTable{
		name.ML: 0.5,
		name.FL: 0.5,
		name.FC: 0.5,
		name.MC: 0.5,
	},
	name.AMR: ZoneEffectTable{
		name.MR: 0.5,
		name.FR: 0.5,
		name.FC: 0.5,
		name.MC: 0.5,
	},
	name.STC: ZoneEffectTable{
		name.FC: 1.0,
		name.FL: 0.5,
		name.FR: 0.5,
	},
}