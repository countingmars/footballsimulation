package foundation

import "sort"

type Ability struct {
	Name string
	Point int
}


type sortable []*Ability

func (array sortable) Len() int {
	return len(array)
}
func (array sortable) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
func (array sortable) Less(i, j int) bool {
	return array[i].Point < array[j].Point
}
func (array sortable) Sort() []*Ability {
	sort.Sort(array)
	return array
}