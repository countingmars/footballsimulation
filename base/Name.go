package base

import "strings"

type Name string

func (this Name) In(line string) bool {
	return strings.Contains(string(this), line)
}
