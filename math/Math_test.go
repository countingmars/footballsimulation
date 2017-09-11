package math

import (
	"testing"
)

func TestMath(t *testing.T)  {
	var i int = -1
	actual := Math.Iabs(i)

	if actual != int(1) {
		t.Error("Math.Abs(-1) should return 1")
	}
}