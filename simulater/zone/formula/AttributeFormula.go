package formula

import "github.com/countingmars/fb/base/stats"

type AttributeFormula struct {
	A *stats.Attribute `json:"attribute"`
	F float32 `json:"attribute_factor"`
}

func (this AttributeFormula) Solve() float32 {
	return float32(this.A.Point) * this.F
}
