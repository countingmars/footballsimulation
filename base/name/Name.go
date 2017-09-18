package name

import "strings"


type Name string
func (this Name) In(line string) bool {
	return strings.Contains(string(this), line)
}
type Names []Name
func (this Names) Contains(name Name) bool {
	index := this.IndexOf(name)
	return index >= 0
}
func (this Names) IndexOf(name Name) int {
	for index, each := range this {
		if each == name {
			return index
		}
	}
	return -1
}


