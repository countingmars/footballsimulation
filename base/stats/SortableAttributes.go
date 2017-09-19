package stats

import "sort"

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