package effect

import "github.com/countingmars/fb/base/name"

type PositionZoneEffect map[name.Name]Effect

func (this PositionZoneEffect) ZoneNamesFor(positionName name.Name) name.Names {
	zoneNames := name.Names{}
	for zoneName, _ := range this[positionName] {
		zoneNames = append(zoneNames, zoneName)
	}
	return zoneNames
}


var PositionZoneEffects = PositionZoneEffect{
	name.GK: {
		name.GK: 2.0,
	},
	name.DC: {
		name.DC: 1.0,
		name.DL: 0.4,
		name.DR: 0.4,
		name.MC: 0.2,
	},
	name.DL: {
		name.DL: 1.0,
		name.ML: 0.5,
		name.DC: 0.5,
	},
	name.DR: {
		name.DR: 1.0,
		name.MR: 0.5,
		name.DC: 0.5,
	},
	name.DMC: {
		name.MC: 0.5,
		name.DC: 0.5,
		name.DL: 0.5,
		name.DR: 0.5,
	},
	name.WBL: {
		name.DL: 0.5,
		name.ML: 0.5,
		name.FL: 0.5,
		name.MC: 0.5,
	},
	name.WBR: {
		name.DR: 0.5,
		name.MR: 0.5,
		name.FR: 0.5,
		name.MC: 0.5,
	},
	name.MC: {
		name.MC: 1.0,
		name.FC: 0.5,
		name.ML: 0.5,
		name.MR: 0.5,
	},
	name.ML: {
		name.ML: 1.0,
		name.DL: 0.5,
		name.FL: 0.5,
		name.MC: 0.5,
	},
	name.AMC: {
		name.MC: 0.5,
		name.FC: 0.5,
		name.FL: 0.5,
		name.FR: 0.5,
	},
	name.AML: {
		name.ML: 0.5,
		name.FL: 0.5,
		name.FC: 0.5,
		name.MC: 0.5,
	},
	name.AMR: {
		name.MR: 0.5,
		name.FR: 0.5,
		name.FC: 0.5,
		name.MC: 0.5,
	},
	name.STC: {
		name.FC: 1.0,
		name.FL: 0.5,
		name.FR: 0.5,
	},
}

