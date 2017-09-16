package simulate

import "github.com/countingmars/fb/core"

type Simulation struct {
	Home   *core.Team
	Away   *core.Team
	First  Highlights
	Second Highlights
}

func (this Simulation) Win() bool {
	if this.Draw() {
		return false
	}
	if winner, _ := this.Winner(); winner == this.Home {
		return true
	} else {
		return false
	}
}
func (this Simulation) Lose() bool {
	if this.Draw() {
		return false
	}
	return !this.Win()
}

func (this Simulation) Draw() bool {
	return this.Goals().Draw()
}
func (this Simulation) Winner() (winner *core.Team, ok bool) {
	if this.Goals().Home > this.Goals().Away {
		winner = this.Home
		ok = true
	} else if this.Goals().Away > this.Goals().Home {
		winner = this.Away
		ok = true
	} else {
		ok = false
	}
	return
}
func (this Simulation) Goals() Goals {
	return Goals{
		Home: this.First.Goals(Left) + this.Second.Goals(Right),
		Away: this.First.Goals(Right) + this.Second.Goals(Left),
	}
}
