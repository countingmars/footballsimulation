package simulater

import (
	"github.com/countingmars/fb/simulater/dice"
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/simulater/zone"
)

func Simulate(home *team.Team, away *team.Team) *Simulation {
	homeEntries := zone.MakeEntries(home.Formation)
	awayEntries := zone.MakeEntries(away.Formation)
	homeZones := zone.ZonesFrom(homeEntries)
	awayZones := zone.ZonesFrom(awayEntries)

	first := simulateHalf(homeZones, awayZones)
	second := simulateHalf(awayZones, homeZones)

	simulation := &Simulation{}
	simulation.Home = home
	simulation.Away = away
	simulation.First = first
	simulation.Second = second
	simulation.HomeEntries = homeEntries
	simulation.AwayEntries = awayEntries
	return simulation
}
func newSituation(left zone.Zones, right zone.Zones) *Situation {
	return &Situation{
		Timer: &Timer{},
		Ball:  &Ball{},
		Side: Left,
		Left:  left,
		Right: right,
	}
}
func simulateHalf(left zone.Zones, right zone.Zones) Highlights {
	var highlights = Highlights{}

	situation := newSituation(left, right)
	situation.KickOff()

	for false == situation.Timer.Over() {
		hl := simulateAttack(situation)
		highlights = append(highlights, hl)

		situation.Side = situation.Side.Reverse()
		situation.Timer.Go()
	}

	return highlights
}
func (this *Situation) KickOff() {
	this.Ball.KickOff()
}
func simulateAttack(situation *Situation) Highlight {
	hl := simulateHighlight(situation)
	hl.Side = situation.Side
	hl.Time = situation.Timer.Now()
	return hl
}

func simulateHighlight(situation *Situation) Highlight {
	possession := 1
	for {
		if false == simulateBuildup(situation) {
			hl := Highlight{}
			hl.Possession = possession
			return hl
		}
		possession++

		if situation.Ball.CanFinish(situation.Side) {
			hl := simulateFinish(situation)
			hl.Possession = possession
			situation.Ball.KickOff()
			return hl
		}
		situation.Ball.Forward(situation.Side)
	}
}

func simulateBuildup(situation *Situation) bool {
	offence := zone.MakeZoneFormula(situation.OffenderZone()).Offence()
	defence := zone.MakeZoneFormula(situation.DefenderZone()).Defence()

	if 0 < dice.Throw(offence) - dice.Throw(defence) {
		situation.OffenderZone().MakeGradeUp()
		situation.DefenderZone().MakeGradeDown()
		return true
	} else {
		situation.DefenderZone().MakeGradeUp()
		situation.OffenderZone().MakeGradeDown()
		return false
	}
}

func simulateFinish(situation *Situation) Highlight {
	scoring := zone.MakeZoneFormula(situation.OffenderZone()).Scoring()
	point := dice.Throw(scoring)

	if point > scoring / 10 {
		return simulateGoal()
	} else {
		return simulateShoot()
	}
}
func simulateGoal() Highlight {
	return pick(GoalHighlights)
}
func simulateShoot() Highlight {
	return pick(ShootHighlights)
}
func pick(highlights []Highlight) Highlight {
	index := dice.Throw(cap(highlights))
	return highlights[index]
}