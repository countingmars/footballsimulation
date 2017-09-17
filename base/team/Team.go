package team

import "github.com/countingmars/fb/base/formation"

type Team struct {
	Name    string
	Formation formation.Formation
}

func (this* Team) Equals(other *Team) bool {
	return this.Name == other.Name
}

func (this *Team) Clone() *Team {
	return &Team{
		Name: this.Name,
		Formation: this.Formation.Clone(),
	}
}