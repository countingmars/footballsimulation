package stats

import (
	"sort"
	"github.com/countingmars/fb/base"
)

type Attribute struct {
	Name base.Name `json:"name"`
	Point int `json:"point"`
}

func (this *Attribute) Clone() Attribute {
	return Attribute{
		this.Name,
		this.Point,
	}
}
type SortableAttributes []*Attribute

func (this SortableAttributes) Len() int {
	return len(this)
}
func (this SortableAttributes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this SortableAttributes) Less(i, j int) bool {
	return this[i].Point < this[j].Point
}
func (this SortableAttributes) Sort() []*Attribute {
	sort.Sort(this)
	return this
}