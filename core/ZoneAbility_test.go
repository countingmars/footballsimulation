package core

import (
	"testing"
)

func TestZoneAbility(t *testing.T) {
	ability := ZoneAbility{
		Attributes: map[string]*Attribute {
			AB_ATTACK: &Attribute{
				Name: AB_ATTACK,
				Point: 1,
			},
			AB_DEFENCE: &Attribute{
				Name: AB_DEFENCE,
				Point: 1,
			},
		},
	}

	actual := ability.Sum()

	if actual != 2 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}