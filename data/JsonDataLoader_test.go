package data

import (
	"testing"
	"github.com/countingmars/fb/core"
)

func TestLoadAbility(t *testing.T) {
	actual := &core.Attribute{}
	LoadJsonFile("json/ability.json", actual)

	if actual.Point != 19 {
		t.Error("speed should be 19, but ", actual.Point)
	}
}
func TestLoadZoneAbility(t *testing.T) {
	actual := &core.ZoneAbility{}
	LoadJsonFile("json/zone-ability.json", actual)

	if actual.Sum() != 38 {
		t.Error("speed should be 19, but ", actual.Sum())
	}
}
func TestLoadTeamAbility(t *testing.T) {
	actual := &core.TeamAbility{}
	LoadJsonFile("json/team-ability.json", actual)

	if actual.Sum() != 38 {
		t.Error("speed should be 38, but ", actual.Sum())
	}
}
func TestLoadTeam(t *testing.T) {
	actual := &core.Team{}
	LoadJsonFile("json/team-perfect.json", actual)

	if 2600 != actual.Ability.Sum() {
		t.Error("This team's power should be 2600, but ", actual.Ability.Sum())
	}
}
