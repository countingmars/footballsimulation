package simulate

type Timer struct {
	time int
}
func (this *Timer) Now() int {
	return this.time
}
func (this *Timer) Go() bool {
	this.time++
	return false == this.Over()
}
func (this *Timer) Over() bool {
	return this.time >= 10
}
