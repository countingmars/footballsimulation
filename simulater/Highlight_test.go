package simulater

import "testing"

func TestHighlights_Append(t *testing.T) {
	hs := Highlights{}

	hs = append(hs, Highlight{})

	if len(hs) != 1 {
		t.Error("failed to append, expected len 1 but ", len(hs))
	}
}
func TestHighlights_Possession(t *testing.T) {
	hls := Highlights{
		Highlight{Possession: 1, Side: Right},
		Highlight{Possession: 2, Side: Right},
	}

	if 3 != hls.Possession(Right) {
		t.Error("expected 3, but ", hls.Possession(Right))
	}
}
func TestHighlights_GoalsShoots(t *testing.T) {
	hls := Highlights{
		Highlight{Possession: 1, Side: Right, Type: Shoot},
		Highlight{Possession: 1, Side: Right, Type: Shoot},
		Highlight{Possession: 2, Side: Right, Type: Goal},
		Highlight{Possession: 2, Side: Right, Type: Goal},
		Highlight{Possession: 2, Side: Right, Type: Goal},
	}

	if 3 != hls.Goals(Right) {
		t.Error("goals expected 3, but ", hls.Goals(Right))
	}
	if 2 != hls.Shoots(Right) {
		t.Error("goals expected 2, but ", hls.Shoots(Right))
	}
}