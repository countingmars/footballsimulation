package stats

import "testing"

func TestAbility_Sum(t *testing.T) {
	ability := Ability{
		Names.Attack: &Attribute{Names.Attack, 1},
		Names.Defence: &Attribute{Names.Defence, 1},
	}

	actual := ability.Sum()

	if actual != 2 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}