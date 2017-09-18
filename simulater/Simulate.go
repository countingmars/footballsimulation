package simulater

import (
	"github.com/countingmars/fb/simulater/dice"
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/simulater/zone"
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/simulater/zone/effect"
)

func Simulate(home *team.Team, away *team.Team) *Simulation {
	homeZones := zone.ZonesFrom(home.Formation)
	awayZones := zone.ZonesFrom(away.Formation)

	first := simulateHalf(homeZones, awayZones)
	second := simulateHalf(awayZones, homeZones)

	simulation := &Simulation{}
	simulation.Home = home
	simulation.Away = away
	simulation.First = first
	simulation.Second = second
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
		situation.Ball.Forward(situation.Side)

		if situation.Ball.CanFinish(situation.Side) {
			hl := simulateFinish(situation)
			hl.Possession = possession
			situation.Ball.KickOff()
			return hl
		}
	}
}
type Play interface {
	Play()
}
type BuildupPlay struct {
	Offender zone.Zones
	Defender zone.Zones
	ZoneName name.Name
}
func (this BuildupPlay) Play() bool {
	offence := int(this.offence(this.Offender[this.ZoneName]))
	defence := int(this.defence(this.Defender[this.ZoneName]))
	if 0 < dice.Throw(offence) - dice.Throw(defence) {
		return true
	} else {
		return false
	}
}
func (this BuildupPlay) offence(zone *zone.Zone) float32 {
	var sum float32
	for _, entry := range zone.Entries {
		playerFormula := effect.PlayerFormula{zone.Name, entry.Position.Name, entry.Player}
		sum += playerFormula.Offence()
	}
	return sum
}
func (this BuildupPlay) defence(zone *zone.Zone) float32 {
	var sum float32
	for _, entry := range zone.Entries {
		playerFormula := effect.PlayerFormula{zone.Name, entry.Position.Name, entry.Player}
		sum += playerFormula.Defence()
	}
	return sum
}
func simulateBuildup(situation *Situation) bool {
	play := BuildupPlay{}
	play.ZoneName = situation.ZoneName()
	play.Offender = situation.Offender()
	play.Defender = situation.Defender()
	return play.Play()
}

func simulateFinish(situation *Situation) Highlight {
	scoring := situation.Offender()[situation.ZoneName()].Scoring()
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