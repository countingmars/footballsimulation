package formation

type Roles []*Role

func (this Roles) Clone() Roles {
	clone := Roles{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}
