package zone

import (
	"testing"
	"github.com/countingmars/fb/base/attr"
)

func TestZoneAbility(t *testing.T) {
	ability := Ability{
		Attributes: map[attr.Name]*attr.Attribute {
			attr.Names.Attack: &attr.Attribute{
				Name: attr.Names.Attack,
				Point: 1,
			},
			attr.Names.Defence: &attr.Attribute{
				Name: attr.Names.Defence,
				Point: 1,
			},
		},
	}

	actual := ability.Sum()

	if actual != 2 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}