package game

import (
	"testing"
	. "github.com/countingmars/fb/foundation"
)

func TestBall_PositionForLeft(t *testing.T) {
	ball := Ball{}

	testPositionFor(t, ball, Left)
	testPositionFor(t, ball, Right)
}

func testPositionFor(t *testing.T, ball Ball, side Side) {
	ball.KickOff(side)
	ballPositionShouldBe(ball.Zone(), "M", t)

	ball.Forward()
	ballPositionShouldBe(ball.Zone(), "F", t)

	ball.Forward()
	if Zones.GK != ball.Zone() {
		t.Error("chance should be returned as true")
	}
}

func ballPositionShouldBe(zone Zone, line string, t *testing.T) {
	if false == zone.In(line) {
		t.Error("ball position expected " + line + ", but ", zone)
	}
}
