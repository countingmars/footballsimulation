package formula

import (

)

type Formula struct {
	Formulas   []AttributeFormula
	ZoneFactor float32 `json:"zone_factor"`
}

func (this *Formula) Solve() float32 {
	var sum float32 = 0
	for _, each := range this.Formulas {
		sum += each.Solve()
	}
	sum *= this.ZoneFactor
	return sum
}

func (this *Formula) Add(af AttributeFormula) {
	this.Formulas = append(this.Formulas, af)
}
