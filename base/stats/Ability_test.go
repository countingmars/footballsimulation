package stats

import (
	"testing"
	"github.com/countingmars/fb/base/name"
)

func TestAbility_Sum(t *testing.T) {
	ability := Ability{
		name.Attack: &Attribute{name.Attack, 1},
		name.Defence: &Attribute{name.Defence, 1},
	}

	actual := ability.Sum()

	if actual != 2 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}