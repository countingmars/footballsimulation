package zone

import "testing"

func TestZone_Ability(t *testing.T) {
	zone := Zone{}

	if nil == zone.Ability() {
		t.Error("ability is null")
	}
}
