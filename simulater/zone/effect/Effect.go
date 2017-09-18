package effect

import "github.com/countingmars/fb/base/name"


type Effect map[name.Name]float32

func (this Effect) Contains(name name.Name) bool {
	_, ok := this[name]
	return ok
}