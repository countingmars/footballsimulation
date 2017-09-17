package position

import "github.com/countingmars/fb/base"


type Position struct {
	Name  base.Name `json:"name"`
	Point int       `json:"point"`
}
type Positions []Position

func (this Positions) Clone() Positions {
	clone := Positions{}
	for _, position := range this {
		clone = append(clone, position)
	}
	return clone
}

func New(name base.Name) Position {
	return Position{Name: name}
}

var (
	GK = Position{Name: Names.GK}
	DL = Position{Name: Names.DL}
	DR = Position{Name: Names.DR}
	DC = Position{Name: Names.DC}
	WBL = Position{Name: Names.WBL}
	WBR = Position{Name: Names.WBR}
	DMC = Position{Name: Names.DMC}
	ML = Position{Name: Names.ML}
	MR = Position{Name: Names.MR}
	MC = Position{Name: Names.MC}
	AML = Position{Name: Names.AML}
	AMR = Position{Name: Names.AMR}
	AMC = Position{Name: Names.AMC}
	STC = Position{Name: Names.STC}
)