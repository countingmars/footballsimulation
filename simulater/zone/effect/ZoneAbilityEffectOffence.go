package effect

import "github.com/countingmars/fb/base/name"

var OffenceZoneAbilityEffect = ZoneAbilityEffect{
	name.DL: {
		name.Pass:    1.5, // technical factors
		name.Dribble: 1.5,
		name.Cross:   3.0,
		name.OffTheBall:   1.5, // mental factors
		name.Decision:     1.2,
		name.Workrate:     2.0,
		name.Anticipation: 1.5,
		name.Speed:    2.0, // physical factors
		name.Attack:   2.0,
		name.Teamwork: 2.0,
	},
	name.DR: {
		name.Pass:    1.5, // technical factors
		name.Dribble: 1.5,
		name.Cross:   3.0,
		name.OffTheBall:   1.5, // mental factors
		name.Decision:     1.2,
		name.Workrate:     2.0,
		name.Anticipation: 1.5,
		name.Speed:    2.0, // physical factors
		name.Attack:   2.0,
		name.Teamwork: 2.0,
	},
	name.DC: {
		name.Pass:    2.0, // technical factors
		name.Dribble: 1.0,
		name.OffTheBall:   1.5, // mental factors
		name.Decision:     1.2,
		name.Workrate:     1.5,
		name.Anticipation: 1.5,
		name.Speed:    1.2, // physical factors
		name.Attack:   1.5,
		name.Teamwork: 2.0,
	},
	name.MC: {
		name.Pass:      3.0, // technical factors
		name.Dribble:   2.0,
		name.Cross:     1.5,
		name.LongShoot: 1.5,
		name.OffTheBall:   2.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.5,
		name.Anticipation: 2.0,
		name.Speed:    1.0, // physical factors
		name.Attack:   2.0,
		name.Teamwork: 3.0,
	},
	name.ML: {
		name.Pass:      2.0, // technical factors
		name.Dribble:   3.0,
		name.Cross:     2.5,
		name.LongShoot: 1.2,
		name.OffTheBall:   2.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.5,
		name.Anticipation: 2.0,
		name.Speed:    3.0, // physical factors
		name.Attack:   2.0,
		name.Teamwork: 1.5,
	},
	name.MR: {
		name.Pass:      2.0, // technical factors
		name.Dribble:   3.0,
		name.Cross:     2.5,
		name.LongShoot: 1.2,
		name.OffTheBall:   2.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.5,
		name.Anticipation: 2.0,
		name.Speed:    3.0, // physical factors
		name.Attack:   2.0,
		name.Teamwork: 1.5,
	},
	name.FC: {
		name.Pass:      2.0, // technical factors
		name.Dribble:   3.0,
		name.Cross:     2.0,
		name.LongShoot: 1.0,
		name.OffTheBall:   3.0, // mental factors
		name.Decision:     3.0,
		name.Workrate:     1.5,
		name.Anticipation: 2.5,
		name.Speed:    2.0, // physical factors
		name.Attack:   3.0,
		name.Teamwork: 1.5,
	},
	name.FL: {
		name.Pass:      2.0, // technical factors
		name.Dribble:   3.0,
		name.Cross:     2.5,
		name.LongShoot: 1.0,
		name.OffTheBall:   3.0, // mental factors
		name.Decision:     2.5,
		name.Workrate:     1.5,
		name.Anticipation: 2.0,
		name.Speed:    2.0, // physical factors
		name.Attack:   3.0,
		name.Teamwork: 1.5,
	},
	name.FR: {
		name.Pass:      2.0, // technical factors
		name.Dribble:   3.0,
		name.Cross:     2.5,
		name.LongShoot: 1.0,
		name.OffTheBall:   3.0, // mental factors
		name.Decision:     2.5,
		name.Workrate:     1.5,
		name.Anticipation: 2.0,
		name.Speed:    2.0, // physical factors
		name.Attack:   3.0,
		name.Teamwork: 1.5,
	},
}

