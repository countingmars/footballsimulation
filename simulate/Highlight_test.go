package simulate

import "testing"

func TestHighlights_Append(t *testing.T) {
	hs := Highlights{}

	hs = append(hs, Highlight{})

	if len(hs) != 1 {
		t.Error("failed to append, expected len 1 but ", len(hs))
	}
}