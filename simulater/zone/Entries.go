package zone

import (
	"github.com/countingmars/fb/base/formation"
	"encoding/json"
)

type Entries []*Entry


func MakeEntries(aFormation formation.Formation) Entries {
	entries := Entries{}
	for _, role := range aFormation {
		entry := &Entry{role.Player, role.Position, Stats{Grade: 6.0}}
		entries = append(entries, entry)
	}
	return entries
}
func (this Entries) Clone() Entries {
	clone := Entries{}
	for _, each := range this {
		clone = append(clone, each.Clone())
	}
	return clone
}
func (this Entries) Json() string {
	out, _ := json.Marshal(this)
	return string(out)
}

func (this Entries) Len() int {
	return len(this)
}

func (this Entries) Less(i, j int) bool {
	return this[i].Stats.LastPoint > this[j].Stats.LastPoint;
}

func (this Entries) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
