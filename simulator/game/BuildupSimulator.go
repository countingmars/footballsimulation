package game

import . "github.com/countingmars/fb/foundation"


type AttackSituation struct {
	Offender TeamAbility
	Defender TeamAbility
	Ball *Ball
	Chance bool
	Success bool
	HighlightSimulator HighlightSimulator
}

func (this *AttackSituation) Play() Highlight {
	buildupSimulator := BuildupSimulator{}
	for {
		buildupSimulator.Simulate(this)
		if this.Success {
			if this.Ball.Zone() == Zones.GK {
				penetration := this.simulatePenetration(this.Offender, this.Defender)
				if penetration.Scored() {
					return this.HighlightSimulator.SimulateGoal()
				} else {
					return this.HighlightSimulator.SimulateShoot()
				}
			}
		} else {
			return Highlight{}
		}
	}
}


func (this *AttackSituation)  simulatePenetration(offender TeamAbility, defender TeamAbility) Penetration {
	scoring := offender[this.Ball.Zone()].Scoring()
	point := Dice(scoring).Throw()

	if point > scoring / 10 {
		return Penetration{ point }
	} else {
		return Penetration{}
	}
}

type BuildupSimulator struct {

}
func (this BuildupSimulator) Simulate(context *AttackSituation) {
	offence := context.Offender.Offence(context.Ball.Zone())
	defence := context.Defender.Defence(context.Ball.Zone())

	if Referee.Judge(offence, defence) {
		context.Ball.Forward()
		context.Success = true
	}
}