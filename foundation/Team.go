package foundation

type Team struct {
	Name    string
	Ability TeamAbility
}

func (t* Team) Equals(other *Team) bool {
	return t.Name == other.Name
}

