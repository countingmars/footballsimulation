package dice

import (
	"testing"
)

func TestHedronDice_Throw(t *testing.T) {
	for {
		point := Throw(3)

		if point == 0 {
			break
		} else if point == 3 {
			t.Error("Dice max should be 2, when Dice(3).Throw()")
		}

	}
}