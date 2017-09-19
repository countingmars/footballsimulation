package zone

import (
	"github.com/countingmars/fb/base/name"

)

type Zone struct {
	Name  name.Name
	Entries Entries
}

func (this *Zone) MakeGradeDown() {
	var num float32 = 0
	for _, entry := range this.Entries {
		entry.Stats.Grade -= num * 0.1
		num++
	}
}

func (this *Zone) MakeGradeUp() {
	num := float32(len(this.Entries))
	for _, entry := range this.Entries {
		entry.Stats.Grade += num * 0.1
		num--
	}
}
func (this *Zone) Add(entry *Entry) {
	this.Entries = append(this.Entries, entry)
}

func (this *Zone) Clone() *Zone {
	return &Zone{
		Name:  this.Name,
		Entries: this.Entries.Clone(),
	}
}

