package attr

import "sort"

type Attribute struct {
	Name Name
	Point int
}


type Attributes []*Attribute

func (this Attributes) Len() int {
	return len(this)
}
func (this Attributes) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this Attributes) Less(i, j int) bool {
	return this[i].Point < this[j].Point
}
func (this Attributes) Sort() []*Attribute {
	sort.Sort(this)
	return this
}