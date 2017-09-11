package foundation

import (
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type HedronDice struct {
	hedron int
}

func Dice(hedron int) HedronDice {
	dice := HedronDice{ hedron }
	return dice
}

func (this HedronDice) Throw() int {
	if this.hedron < 1 {
		return 0
	}
	return rand.Intn(this.hedron)
}
func (this HedronDice) Throws(times int) []int{
	numbers := make([]int, times)
	for i := 0; i < times; i++ {
		numbers[i] = this.Throw()
	}
	return numbers
}
