package foundation

import (
	"testing"
)


func newZoneAbility(name string, point int) ZoneAbility {
	return ZoneAbility{
			name: &Ability{name, point},
		}
}
func TestTeamAbility_Sum(t *testing.T) {
	teamAbility := TeamAbility{}
	teamAbility[DL] = newZoneAbility(AB_ATTACK, 1)
	teamAbility[DR] = newZoneAbility(AB_DEFENCE, 1)

	actual := teamAbility.Sum()

	if 2 != actual {
		t.Error("The sum of the power expected 2, but ", actual)
	}
}

func TestTeamAbility_Inspect(t *testing.T) {
	teamAbility := TeamAbility{}
	teamAbility[DL] = newZoneAbility(AB_ATTACK, 3)
	teamAbility[ML] = newZoneAbility(AB_ATTACK, 6)
	teamAbility[FL] = newZoneAbility(AB_ATTACK, 9)

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
func TestTeamAbility_Effect(t *testing.T) {
	ability := TeamAbility{}
	ability[DL] = newZoneAbility(AB_ATTACK, 150)
	ability[DC] = newZoneAbility(AB_ATTACK, 200)

	actual := Effect{ability[DL], ability[DC], 0.35}.Calculate()

	if actual != 17 {
		t.Error("The effect of DL to DC should be 17, but ", actual)
	}
}
