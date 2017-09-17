package player

import (
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base"
)

func TestPlayer(aPosition position.Position) *Player {
	player := &Player{}
	player.Name = "player-name"
	player.Age = 21
	player.Country = base.Countries.DE
	player.Positions = []position.Position{aPosition}
	player.Ability = stats.Ability{
		stats.Names.Speed: {Name: stats.Names.Speed, Point: 20},
		stats.Names.Health: {Name: stats.Names.Health, Point: 20},
		stats.Names.Workrate: {Name: stats.Names.Workrate, Point: 20},
		stats.Names.Strength: {Name: stats.Names.Strength, Point: 20},

		// technical
		stats.Names.Dribble: {Name: stats.Names.Dribble, Point: 20},
		stats.Names.Pass: {Name: stats.Names.Pass, Point: 20},
		stats.Names.Cross: {Name: stats.Names.Cross, Point: 20},
		stats.Names.Shoot: {Name: stats.Names.Shoot, Point: 20},
		stats.Names.LongShoot: {Name: stats.Names.LongShoot, Point: 20},
		stats.Names.Tackle: {Name: stats.Names.Tackle, Point: 20},
		stats.Names.Mark: {Name: stats.Names.Mark, Point: 20},
		stats.Names.Finish: {Name: stats.Names.Finish, Point: 20},
		stats.Names.Heading: {Name: stats.Names.Heading, Point: 20},

		// mental
		stats.Names.Anticipation: {Name: stats.Names.Anticipation, Point: 20},
		stats.Names.Creative: {Name: stats.Names.Creative, Point: 20},
		stats.Names.Positioning: {Name: stats.Names.Positioning, Point: 20},
		stats.Names.OffTheBall: {Name: stats.Names.OffTheBall, Point: 20},
		stats.Names.Teamwork: {Name: stats.Names.Teamwork, Point: 20},
		stats.Names.Decision: {Name: stats.Names.Decision, Point: 20},
		stats.Names.Strategy: {Name: stats.Names.Strategy, Point: 20},

		stats.Names.Defence: {Name: stats.Names.Defence, Point: 20},
		stats.Names.Attack: {Name: stats.Names.Attack, Point: 20},
		stats.Names.Gk: {Name: stats.Names.Gk, Point: 20},
	}
	return player
}

