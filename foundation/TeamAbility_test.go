package foundation

import (
	"testing"
)

func newZoneAbility(zoneName string, attrName string, point int) ZoneAbility {
	return ZoneAbility{
		Name: zoneName,
		Attributes: map[string]*Attribute { attrName: &Attribute{attrName, point }},
	}
}
func TestTeamAbility_Sum(t *testing.T) {
	teamAbility := TeamAbility{}
	teamAbility[DL] = newZoneAbility(DL, AB_ATTACK, 1)
	teamAbility[DR] = newZoneAbility(DR, AB_DEFENCE, 1)

	actual := teamAbility.Sum()

	if 2 != actual {
		t.Error("The sum of the power expected 2, but ", actual)
	}
}

func TestTeamAbility_Inspect(t *testing.T) {
	teamAbility := TeamAbility{}
	teamAbility[DL] = newZoneAbility(DL, AB_ATTACK, 3)
	teamAbility[ML] = newZoneAbility(ML, AB_ATTACK, 6)
	teamAbility[FL] = newZoneAbility(FL, AB_ATTACK, 9)

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
	teamAbility := TeamAbility{
		MC: newZoneAbility(MC, AB_PASS, 10),
	}

	actual := teamAbility.Possession()

	if actual != 60 {
		t.Error("expected possession 60 (pass * attr factory 30 * zone factor 2), but ", actual)
	}
}