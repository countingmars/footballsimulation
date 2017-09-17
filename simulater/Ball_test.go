package simulater

import (
	"testing"
	"github.com/countingmars/fb/base/zone"
	"github.com/countingmars/fb/base"
)

func TestBall_Zone(t *testing.T) {
	ball := Ball{}
	ball.KickOff()

	if zone.Names.MC != ball.ZoneName(Left) {
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

	ball.Forward(side)
	if zone.Names.GK != ball.ZoneName(side) {
		t.Error("chance should be returned as true")
	}
}

func ballPositionShouldBe(zone base.Name, line string, t *testing.T) {
	if false == zone.In(line) {
		t.Error("ball position expected " + line + ", but ", zone)
	}
}
