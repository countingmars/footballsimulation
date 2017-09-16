package game

type Timer struct {
	time int
}
func (this *Timer) Now() int {
	return this.time
}
func (this *Timer) Go() int {
	this.time++
	return this.time
}
func (this *Timer) Over() bool {
	return this.time >= 10
}
