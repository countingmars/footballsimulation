package player

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/base/name"
)

func TestPlayer(aPosition position.Position) *Player {
	player := &Player{}
	player.Name = "player-name"
	player.Age = 21
	player.Country = base.Countries.DE
	player.Positions = []position.Position{aPosition}
	player.Ability = stats.Ability{
		name.Speed: {Name: name.Speed, Point: 20},
		name.Health: {Name: name.Health, Point: 20},
		name.Workrate: {Name: name.Workrate, Point: 20},
		name.Strength: {Name: name.Strength, Point: 20},

		// technical
		name.Dribble: {Name: name.Dribble, Point: 20},
		name.Pass: {Name: name.Pass, Point: 20},
		name.Cross: {Name: name.Cross, Point: 20},
		name.Shoot: {Name: name.Shoot, Point: 20},
		name.LongShoot: {Name: name.LongShoot, Point: 20},
		name.Tackle: {Name: name.Tackle, Point: 20},
		name.Mark: {Name: name.Mark, Point: 20},
		name.Finish: {Name: name.Finish, Point: 20},
		name.Heading: {Name: name.Heading, Point: 20},

		// mental
		name.Anticipation: {Name: name.Anticipation, Point: 20},
		name.Creative: {Name: name.Creative, Point: 20},
		name.Positioning: {Name: name.Positioning, Point: 20},
		name.OffTheBall: {Name: name.OffTheBall, Point: 20},
		name.Teamwork: {Name: name.Teamwork, Point: 20},
		name.Decision: {Name: name.Decision, Point: 20},
		name.Strategy: {Name: name.Strategy, Point: 20},

		name.Defence: {Name: name.Defence, Point: 20},
		name.Attack: {Name: name.Attack, Point: 20},
		name.Goalkeeping: {Name: name.Goalkeeping, Point: 20},
	}
	return player
}

