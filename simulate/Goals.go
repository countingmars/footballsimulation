package simulate

type Goals struct {
	Home int
	Away int
}
func (this Goals) Sum() int {
	return this.Home + this.Away
}
func (this Goals) Draw() bool {
	return this.Home == this.Away
}