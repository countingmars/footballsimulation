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

func (this *Game) Amend() *Game {
	homeAbility := this.Home.Ability.AmendBySelf()
	awayAbility := this.Away.Ability.AmendBySelf()

	homeAbilityAmendedByAway := homeAbility.AmendBy(awayAbility)
	awayAbilityAmendedByHome := awayAbility.AmendBy(homeAbility)
	
	return &Game{
		Home: &Team{
			Name: this.Home.Name,
			Ability: homeAbilityAmendedByAway,
		},
		Away: &Team{
			Name: this.Away.Name,
			Ability: awayAbilityAmendedByHome,
		},
	}
}