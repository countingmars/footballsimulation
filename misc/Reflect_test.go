package misc

import (
	"testing"
)


type pair struct {
	A int
	B int
}
func TestClone(t *testing.T) {
	testCloneValues(t)
	testClonePointers(t)
	testCloneValuePointer(t)
}
func TestCloneMaps(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
	}
	m2 := map[string]int{}

	err := Clone(m1, m2)
	if err != nil {
		t.Error("clonning to ptr should be possible")
	}

	_, ok := m2["a"]
	if !ok {
		t.Error("failed to clone map", ok)
	}

	m1["a"] = 2
	if m2["a"] == 2 {
		t.Error("clonning should not copy ptr itself, but only value")
	}
}
func testCloneValuePointer(t *testing.T) {
	value := pair{1,2}
	ptr := &pair{}

	err := Clone(value, ptr)

	if err != nil {
		t.Error("clonning to ptr should be possible")
	}
	if ptr.A != 1 {
		t.Error("p2.A should be 1, but ", ptr.A)
	}
	value.A = 2
	if ptr.A == 2 {
		t.Error("clonning should not copy ptr itself, but only value")
	}
}
func testCloneValues(t *testing.T) {
	s1 := pair{1,2}
	s2 := pair{}
	err := Clone(s1, s2)
	if err == nil {
		t.Error("clonning to value, not ptr, should be impossible")
	}
}
func testClonePointers(t *testing.T) {
	p1 := &pair{1,2}
	p2 := &pair{}

	err := Clone(p1, p2)

	if err != nil {
		t.Error("clonning to ptr should be possible")
	}
	if p2.A != 1 {
		t.Error("p2.A should be 1, but ", p2.A)
	}
	p1.A = 2
	if p2.A == 2 {
		t.Error("clonning should not copy ptr itself, but only value")
	}
}
