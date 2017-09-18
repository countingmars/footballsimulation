package zone

import (
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base/name"
)

type Zone struct {
	Name              name.Name
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

func (this *Zone) calculateAttr(key name.Name, factor float32) float32 {
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



func resolveZoneNames(aPosition position.Position) name.Names {
	return PET.ZoneNamesFor(aPosition.Name)
	/*switch aPosition.Name {
	case name.GK: return name.Names{name.GK}
	case name.DC: return name.Names{name.DC, name.MC, name.DL, name.DR}
	case name.DL: return name.Names{name.DL, name.ML, name.DC}
	case name.DR: return name.Names{name.DR, name.MR, name.DC}
	case name.WBL: return name.Names{name.DL, name.ML, name.FL, name.DC, name.MC}
	case name.WBR: return name.Names{name.DR, name.MR, name.FR, name.DC, name.MC}
	case name.DMC: return name.Names{name.DC, name.MC, name.ML, name.MR, name.DL, name.DR}
	case name.MC: return name.Names{name.DC, name.MC, name.ML, name.MR, name.FC}
	case name.ML: return name.Names{name.DL, name.ML, name.MC, name.FL}
	case name.MR: return name.Names{name.DR, name.MR, name.MC, name.FR}
	case name.AMC: return name.Names{name.MC, name.ML, name.MR, name.FC, name.FR, name.FL}
	case name.AML: return name.Names{name.ML, name.DL, name.FC, name.FL}
	case name.AMR: return name.Names{name.MR, name.DR, name.FC, name.FR}
	case name.STC: return name.Names{name.MC, name.FL, name.FC, name.FR}
	}
	panic("unknown position: " + aPosition.Name)*/
}