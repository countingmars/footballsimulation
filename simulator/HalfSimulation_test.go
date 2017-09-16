package simulator

import "testing"

func TestHalfSimulation_Append(t *testing.T) {
	hs := HalfSimulation{}

	hs = append(hs, Highlight{})

	if len(hs) != 1 {
		t.Error("failed to append, expected len 1 but ", len(hs))
	}
}