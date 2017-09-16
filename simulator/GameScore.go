package simulator

type GameScore struct {
	Home int
	Away int
}
func (score *GameScore) Sum() int {
	return score.Home + score.Away
}
