package math

import "math"

type nothing struct {

}
var Math nothing = nothing{}

func (nothing nothing) Iabs(i int) int {
	if 0 < i {
		return i
	} else {
		return -i
	}
}
func (nothing nothing) Fabs(f float64) float64 {
	return math.Abs(f)
}