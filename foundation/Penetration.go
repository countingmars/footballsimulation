package foundation

type Penetration struct {
	Difficulty int
}

func (this Penetration) Scored() bool {
	return 0 < this.Difficulty
}