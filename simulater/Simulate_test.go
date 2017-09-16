package simulater

import (
	"testing"
	"github.com/countingmars/fb/data"
)

func TestSimulate_simulateBuildup(t *testing.T) {
	home := data.LoadTeam("../data/json/team-perfect.json")
	away := home.Clone()
	away.Ability.Set(0)

	for i := 0; i < 100; i++ {
		if false == simulateBuildup(newSituation(home.Ability, away.Ability)) {
			t.Error("away ability is 0 so should always win")
		}
	}
}

func TestSimulate_simulateAttack(t *testing.T) {
	home := data.LoadTeam("../data/json/team-perfect.json")
	away := home.Clone()
	away.Ability.Set(0)

	hls := Highlights{}
	for i := 0; i < 5; i++ {
		hl := simulateAttack(newSituation(home.Ability, away.Ability))
		hls = append(hls, hl)
	}

	if 0 != hls.Possession(Right) {
		t.Error("ability 0 team's possession should be 0, but ", hls.Possession(Right))
	}
	if 0 == hls.Possession(Left) {
		t.Error("ability 20 team's possession should be greater than 0, but ", hls.Possession(Left))
	}
}