package stats

import (
	"testing"
	"github.com/countingmars/fb/base/name"
)

func TestAttributes_Sum(t *testing.T) {
	attributes := Attributes{
		name.Attack: &Attribute{name.Attack, 1},
		name.Defence: &Attribute{name.Defence, 1},
	}

	actual := attributes.Sum()

	if actual != 2 {
		t.Error("zone ability sum should be 1, but ", actual)
	}
}