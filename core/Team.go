package core

type Team struct {
	Name    string
	Ability TeamAbility
}

func (this* Team) Equals(other *Team) bool {
	return this.Name == other.Name
}

func (this *Team) Clone() *Team {
	return &Team{
		Name: this.Name,
		Ability: this.Ability.Clone(),
	}
}