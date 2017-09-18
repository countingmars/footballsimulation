package simulater

import (
	"testing"
	"github.com/countingmars/fb/base/formation"
	"github.com/countingmars/fb/simulater/zone"
)

func TestSimulate_simulateBuildup(t *testing.T) {
	home := zone.ZonesFrom(formation.Test442())
	away := home.Clone()
	away.Set(0)

	for i := 0; i < 100; i++ {
		if false == simulateBuildup(newSituation(home, away)) {
			t.Error("away ability is 0 so should always win")
			return
		}
	}
}

func TestSimulate_simulateAttack(t *testing.T) {
	home := zone.ZonesFrom(formation.Test442())
	away := home.Clone()
	away.Set(0)

	hls := Highlights{}
	for i := 0; i < 5; i++ {
		hl := simulateAttack(newSituation(home, away))
		hls = append(hls, hl)
	}

	if 0 != hls.Possession(Right) {
		t.Error("ability 0 team's possession should be 0, but ", hls.Possession(Right))
	}
	if 0 == hls.Possession(Left) {
		t.Error("ability 20 team's possession should be greater than 0, but ", hls.Possession(Left))
	}
}
