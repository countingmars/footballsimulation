package dice

import (
	"time"
	"math/rand"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}


func Throw(hedron int) int {
	if hedron < 1 {
		return 0
	}
	return rand.Intn(hedron)
}
func Throws(hedron, times int) []int {
	numbers := make([]int, times)
	for i := 0; i < times; i++ {
		numbers[i] = Throw(hedron)
	}
	return numbers
}
func Judge(a, b int) bool {
	number := Throw(a + b)
	return number < a
}
