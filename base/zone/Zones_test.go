package zone

import (
	"testing"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/formation"
)

func TestZones_Set(t *testing.T) {
	dc := TestZone(
		Names.DC,
		formation.TestPositionedPlayer(position.GK),
		formation.TestPositionedPlayer(position.DC),
	)
	dl := TestZone(
		Names.DL,
		formation.TestPositionedPlayer(position.DL),
		formation.TestPositionedPlayer(position.WBL),
	)
	zones := Zones{dc.Name: &dc, dl.Name: &dl}
	zones.Set(1)

	if 46 != zones.Sum() {
		t.Error("sum should be 46, but ", zones.Sum())
	}
}
