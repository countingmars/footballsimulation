package stats

import (
	"github.com/countingmars/fb/base/name"
)

type Attribute struct {
	Name name.Name `json:"name"`
	Point int `json:"point"`
}

func (this *Attribute) Clone() Attribute {
	return Attribute{
		this.Name,
		this.Point,
	}
}