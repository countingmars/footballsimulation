package formula

import (
	"testing"
	"github.com/countingmars/fb/base/stats"
)

func TestFormula_Solve(t *testing.T) {
	f := Formula{[]AttributeFormula{}, 0.8}
	if 0 != f.Solve() {
		t.Error("empty formula should be 0")
	}

	f.Add(AttributeFormula{&stats.Attribute{"", 20}, 1.0})
	f.Add(AttributeFormula{&stats.Attribute{"", 20}, 1.0})

	if 32 != f.Solve() {
		t.Error("formula.Solve(), expected 32.0, but ", f.Solve())
	}
}
func TestAttrFormula_Solve(t *testing.T) {
	af := AttributeFormula{&stats.Attribute{"", 20}, 1.0}

	actual := af.Solve()

	if actual != 20 {
		t.Error("expected 20, but ", actual)
	}
}
