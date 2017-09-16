package foundation

var Referee = &referee{}

type referee struct {

}

func (this referee) Judge(a int, b int) bool {
	number := Dice(a + b).Throw()
	return number < a
}
