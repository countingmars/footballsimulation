package gamesimulator

import (
	"testing"
	"strings"
)

func TestBall_PositionForLeft(t *testing.T) {
	ball := Ball{}

	testPositionFor(t, ball, Left)
	testPositionFor(t, ball, Right)
}

func testPositionFor(t *testing.T, ball Ball, side Side) {
	ball.KickOff()
	ballPositionShouldBe(ball.Position(side), "M", t)

	ball.Forward(side)
	ballPositionShouldBe(ball.Position(side), "F", t)

	ball.Forward(side)
	if false == ball.Chance(side) {
		t.Error("chance should be returned as true")
	}
}

func ballPositionShouldBe(ballPosition string, line string, t *testing.T) {
	if false == strings.Contains(ballPosition, line) {
		t.Error("ball position expected " + line + ", but ", ballPosition)
	}
}
