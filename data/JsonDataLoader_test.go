package data

import (
	"testing"
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base/player"
)

func TestLoadJsonFile_attribute(t *testing.T) {
	actual := &stats.Attribute{}
	LoadJsonFile("/json/attribute.json", actual)

	if actual.Point != 19 {
		t.Error("speed should be 19, but ", actual.Point)
	}
}
func TestLoadJsonFile_player(t *testing.T) {
	actual := player.Player{}
	LoadJsonFile("/json/player-perfect.json", &actual)

	if actual.Ability.Sum() != 460 {
		t.Error("sum should be 460, but ", actual.Ability.Sum())
	}
}
