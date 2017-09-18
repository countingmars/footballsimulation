package formation

import "encoding/json"

type Roles []*Role

func (this Roles) Clone() Roles {
	clone := Roles{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}
func (this Roles) Json() string {
	out, _ := json.Marshal(this)
	return string(out)
}