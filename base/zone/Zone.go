package zone

import (
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/base/stats"
)

type Zone struct {
	Name              base.Name
	PositionedPlayers formation.PositionedPlayers

	ability           stats.Ability
}
func (this *Zone) Add(player *formation.PositionedPlayer) {
	this.PositionedPlayers = append(this.PositionedPlayers, player)
}
func (this *Zone) Ability() stats.Ability {
	if len(this.ability) > 0 {
		return this.ability
	}
	this.ability = stats.Ability{}
	for _, player := range this.PositionedPlayers {
		this.ability.Add(player.Player.Ability)
	}
	return this.ability
}

func (this *Zone) calculateAttr(key base.Name, factor float32) float32 {
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
		Name: this.Name,
		PositionedPlayers: this.PositionedPlayers.Clone(),
	}
}

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

func resolveZoneNames(aPosition position.Position) []base.Name {
	switch aPosition.Name {
	case position.Names.GK: return []base.Name{Names.GK}
	case position.Names.DC: return []base.Name{Names.DC, Names.MC}
	case position.Names.DL: return []base.Name{Names.DL, Names.ML}
	case position.Names.DR: return []base.Name{Names.DR, Names.MR}
	case position.Names.WBL: return []base.Name{Names.DL, Names.ML, Names.FL}
	case position.Names.WBR: return []base.Name{Names.DR, Names.MR, Names.FR}
	case position.Names.DMC: return []base.Name{Names.DC, Names.MC}
	case position.Names.MC: return []base.Name{Names.DC, Names.MC, Names.ML, Names.MR, Names.FC}
	case position.Names.ML: return []base.Name{Names.DL, Names.ML, Names.MC, Names.FL}
	case position.Names.MR: return []base.Name{Names.DR, Names.MR, Names.MC, Names.FR}
	case position.Names.AMC: return []base.Name{Names.MC, Names.FC, Names.FR, Names.FL}
	case position.Names.AML: return []base.Name{Names.ML, Names.MC, Names.FC, Names.FL}
	case position.Names.AMR: return []base.Name{Names.MR, Names.MC, Names.FC, Names.FR}
	case position.Names.STC: return []base.Name{Names.FL, Names.FC, Names.FR}
	}
	panic("unknown position: " + aPosition.Name)
}