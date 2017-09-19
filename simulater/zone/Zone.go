package zone

import (
	"github.com/countingmars/fb/base/name"
)

type Zone struct {
	Name  name.Name
	Entries Entries
}
func (this *Zone) Add(role *Entry) {
	this.Entries = append(this.Entries, role)
}

func (this *Zone) Sum() int {
	sum := 0
	for _, v := range this.Entries {
		sum += v.Player.Attributes.Sum()
	}
	return sum
}
func (this *Zone) Clone() *Zone {
	return &Zone{
		Name:  this.Name,
		Entries: this.Entries.Clone(),
	}
}

