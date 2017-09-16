package base

import "sort"

type Attribute struct {
	Name string
	Point int
}


type sortable []*Attribute

func (array sortable) Len() int {
	return len(array)
}
func (array sortable) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
func (array sortable) Less(i, j int) bool {
	return array[i].Point < array[j].Point
}
func (array sortable) Sort() []*Attribute {
	sort.Sort(array)
	return array
}