package formation

import (
	"github.com/countingmars/fb/base/position"
)

var F442 = Formation {
	NewRole(position.GK),
	NewRole(position.DL),
	NewRole(position.DR),
	NewRole(position.DC),
	NewRole(position.DC),
	NewRole(position.ML),
	NewRole(position.MR),
	NewRole(position.MC),
	NewRole(position.MC),
	NewRole(position.STC),
	NewRole(position.STC),
}

type Formation []*Role

func (this Formation) Sum() int {
	sum := 0
	for _, role := range this {
		sum += role.Player.Attributes.Sum()
	}
	return sum
}

func (this Formation) Set(point int) {
	for _, role := range this {
		role.Player.Attributes.Set(point)
	}
}

func (this Formation) Clone() Formation {
	clone := Formation{}
	for _, role := range this {
		clone = append(clone, role.Clone())
	}
	return clone
}
