package zone

import (
	"testing"
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/base/name"
	"fmt"
)

func TestZones_ZonesFrom(t *testing.T) {
	zones := ZonesFrom(MakeEntries(formation.F442))

	if 10 != len(zones) {
		t.Error("expected zone count is 10, but ", len(zones))
	}
	assertThatZoneHasPlayers(t, zones[name.GK], 1)
	assertThatZoneHasPlayers(t, zones[name.DL], 4)
	assertThatZoneHasPlayers(t, zones[name.ML], 4)
	assertThatZoneHasPlayers(t, zones[name.FL], 3)
	assertThatZoneHasPlayers(t, zones[name.MC], 5)
}
func assertThatZoneHasPlayers(t *testing.T, aZone *Zone, expected int) {
	if expected == len(aZone.Entries) {
		return
	}

	message := fmt.Sprintf("in 442, %s zone expected players are %d, but %d : %s",
		aZone.Name,
		expected,
		len(aZone.Entries),
		aZone.Entries.Json(),
	)
	panic(message)
}
