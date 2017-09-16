package team

import (
	"testing"
	"github.com/countingmars/fb/base/attr"
	"github.com/countingmars/fb/base/zone"
)

func newZoneAbility(zoneName zone.Name, attrName attr.Name, point int) zone.Ability {
	return zone.Ability{
		Name: zoneName,
		Attributes: map[attr.Name]*attr.Attribute {
			attrName: &attr.Attribute{attrName, point },
		},
	}
}
func TestTeamAbility_Sum(t *testing.T) {
	teamAbility := Ability{}
	teamAbility[zone.Names.DL] = newZoneAbility(zone.Names.DL, attr.Names.Attack, 1)
	teamAbility[zone.Names.DR] = newZoneAbility(zone.Names.DR, attr.Names.Defence, 1)

	actual := teamAbility.Sum()

	if 2 != actual {
		t.Error("The sum of the power expected 2, but ", actual)
	}
}

func TestTeamAbility_Inspect(t *testing.T) {
	teamAbility := Ability{}
	teamAbility[zone.Names.DL] = newZoneAbility(zone.Names.DL, attr.Names.Attack, 3)
	teamAbility[zone.Names.ML] = newZoneAbility(zone.Names.ML, attr.Names.Attack, 6)
	teamAbility[zone.Names.FL] = newZoneAbility(zone.Names.FL, attr.Names.Attack, 9)

	actual := teamAbility.Inspect()

	if actual.Diff() != 2 {
		t.Error("The diff should be 2, but ", actual.Diff())
	}
	if actual.Strength.Name != "F" {
		t.Error("The strength should be Forward line, but", actual.Strength)
	}
	if actual.Weakness.Name != "D" {
		t.Error("The weakness should be Defence line, but", actual.Weakness)
	}
	if actual.Style != TAI_BALANCE {
		t.Error("The style should be balance style, but", actual.Weakness)
	}
}

func TestTeamAbility_Possession(t *testing.T) {
	teamAbility := Ability{
		zone.Names.MC: newZoneAbility(zone.Names.MC, attr.Names.Pass, 10),
	}

	actual := teamAbility.Possession()

	if actual != 60 {
		t.Error("expected possession 60 (pass * attr factory 30 * zone factor 2), but ", actual)
	}
}