package simulater

import (
	"testing"
	"github.com/countingmars/fb/base/name"
)

func TestBall_Zone(t *testing.T) {
	ball := Ball{}
	ball.KickOff()

	if name.MC != ball.ZoneName(Left) {
		t.Error("initial ball position should be MC, but ", ball.ZoneName(Left))
	}
}
func TestBall_PositionForLeft(t *testing.T) {
	ball := Ball{}

	testPositionFor(t, ball, Left)
	testPositionFor(t, ball, Right)
}

func testPositionFor(t *testing.T, ball Ball, side Side) {
	ball.KickOff()
	ballPositionShouldBe(ball.ZoneName(side), "M", t)

	ball.Forward(side)
	ballPositionShouldBe(ball.ZoneName(side), "F", t)

	if false == ball.CanFinish(side) {
		t.Error(ball.ZoneName(side), " position should be finishing position")
	}
}

func ballPositionShouldBe(zone name.Name, line string, t *testing.T) {
	if false == zone.In(line) {
		t.Error("ball position expected " + line + ", but ", zone)
	}
}
