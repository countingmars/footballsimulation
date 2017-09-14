package gamesimulator


type Side string

const (
	Left Side = "left"
	Right Side = "right"
)

func (this Side) Reverse() Side {
	if this == Left {
		return Right
	} else {
		return Left
	}
}
