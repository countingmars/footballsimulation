package position

import (
	"github.com/countingmars/fb/base/name"
	"encoding/json"
)


type Position struct {
	Name  name.Name `json:"name"`
	Point int       `json:"point"`
}
func (this Position) Json() string {
	out, _ := json.Marshal(&this)
	return string(out)
}


type Positions []Position

func (this Positions) Clone() Positions {
	clone := Positions{}
	for _, position := range this {
		clone = append(clone, position)
	}
	return clone
}


func New(name name.Name) Position {
	return Position{Name: name}
}

var (
	GK = Position{Name: name.GK}
	DL = Position{Name: name.DL}
	DR = Position{Name: name.DR}
	DC = Position{Name: name.DC}
	WBL = Position{Name: name.WBL}
	WBR = Position{Name: name.WBR}
	DMC = Position{Name: name.DMC}
	ML = Position{Name: name.ML}
	MR = Position{Name: name.MR}
	MC = Position{Name: name.MC}
	AML = Position{Name: name.AML}
	AMR = Position{Name: name.AMR}
	AMC = Position{Name: name.AMC}
	STC = Position{Name: name.STC}
)