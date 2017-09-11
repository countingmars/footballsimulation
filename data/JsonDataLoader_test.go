package data

import (
	"testing"
	. "github.com/countingmars/fb/foundation"
)

func TestLoadAbility(t *testing.T) {
	actual := &Ability{}
	LoadJsonFile("json/ability.json", actual)

	if actual.Point != 19 {
		t.Error("speed should be 19, but ", actual.Point)
	}
}
func TestLoadZoneAbility(t *testing.T) {
	actual := &ZoneAbility{}
	LoadJsonFile("json/zone-ability.json", actual)

	if actual.Sum() != 38 {
		t.Error("speed should be 19, but ", actual.Sum())
	}
}
func TestLoadTeamAbility(t *testing.T) {
	actual := &TeamAbility{}
	LoadJsonFile("json/team-ability.json", actual)

	if actual.Sum() != 38 {
		t.Error("speed should be 38, but ", actual.Sum())
	}
}
func TestLoadTeam(t *testing.T) {
	actual := &Team{}
	LoadJsonFile("json/team-perfect.json", actual)

	if 1800 != actual.Ability.Sum() {
		t.Error("This team's power should be 1800, but ", actual.Ability.Sum())
	}
}
