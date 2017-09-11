package foundation

import (
	"testing"
)

func TestZoneAbility(t *testing.T) {
	ability := ZoneAbility{
		AB_ATTACK: &Ability{AB_ATTACK, 1},
	}

	actual := ability.Sum()

	if actual != 1 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}
func TestZoneAbility_Debuff(t *testing.T) {
	actual := ZoneAbility{
		AB_ATTACK: &Ability{AB_ATTACK, 2},
		AB_DEFENCE: &Ability{AB_DEFENCE, 2},
		AB_SPEED: &Ability{AB_SPEED, 2},
		AB_HEALTH: &Ability{AB_HEALTH, 2},
	}

	actual.Debuff(4)

	if actual.Sum() != 4 {
		t.Error("debuffed sum expected 4, but ", actual.Sum())
	}
}