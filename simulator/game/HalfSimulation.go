package game

type HalfSimulation []Highlight

func (this HalfSimulation) Goals(side Side) int {
	goals := 0
	for _, v := range this {
		if v.Side == side && v.Type == Goal {
			goals++
		}
	}
	return goals
}
