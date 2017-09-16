package simulater

type Finish int

func (this Finish) Scored() bool {
	return 0 < this
}