package foundation

type Game struct {
	Home *Team
	Away *Team
}
func (this *Game) Sum() int {
	return this.Home.Ability.Sum() + this.Away.Ability.Sum()
}
func (this *Game) AwayAbility() int {
	return this.Away.Ability.Sum()
}
func (this *Game) HomeAbility() int {
	return this.Home.Ability.Sum()
}

