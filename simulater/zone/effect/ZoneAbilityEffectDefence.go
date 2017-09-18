package effect

import "github.com/countingmars/fb/base/name"

var DefenceZoneAbilityEffect = ZoneAbilityEffect{
	name.DL: {
		name.Tackle: 2.5, // tech
		name.Mark: 3.0,
		name.Positioning: 2.5, // mental
		name.Decision: 2.0,
		name.Workrate: 2.0,
		name.Anticipation: 2.0,
		name.Teamwork: 2.0,
		name.Defence: 2.0,
		name.Speed: 2.0, // physical
		name.Strength: 1.5,
	},
	name.DR: {
		name.Tackle: 2.5, // tech
		name.Mark: 3.0,
		name.Positioning: 2.5, // mental
		name.Decision: 2.0,
		name.Workrate: 2.0,
		name.Anticipation: 2.0,
		name.Teamwork: 2.0,
		name.Defence: 2.0,
		name.Speed: 2.0, // physical
		name.Strength: 1.5,
	},
	name.DC: {
		name.Tackle: 2.0, // technical factors
		name.Mark:   3.0,
		name.Positioning:  3.0, // mental factors
		name.Decision:     2.0,
		name.Workrate:     1.5,
		name.Anticipation: 2.0,
		name.Teamwork:     3.0,
		name.Defence:      3.0,
		name.Speed:    1.2, // physical factors
		name.Strength: 3.0,
	},
	name.MC: {
		name.Tackle: 1.5, // technical factors
		name.Mark:   2.0,
		name.Positioning:  2.5, // mental factors
		name.Decision:     2.0,
		name.Workrate:     2.0,
		name.Anticipation: 2.0,
		name.Teamwork:     2.5,
		name.Defence:      1.5,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
	name.ML: {
		name.Tackle: 1.0, // technical factors
		name.Mark:   1.5,
		name.Positioning:  1.5, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.0,
		name.Anticipation: 2.0,
		name.Teamwork:     1.5,
		name.Defence:      1.0,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
	name.MR: {
		name.Tackle: 1.0, // technical factors
		name.Mark:   1.5,
		name.Positioning:  1.5, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.0,
		name.Anticipation: 2.0,
		name.Teamwork:     1.5,
		name.Defence:      1.0,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
	name.FC: {
		name.Tackle: 1.0, // technical factors
		name.Mark:   1.0,
		name.Positioning:  1.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.0,
		name.Anticipation: 1.5,
		name.Teamwork:     1.5,
		name.Defence:      0.8,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
	name.FL: {
		name.Tackle: 1.0, // technical factors
		name.Mark:   1.0,
		name.Positioning:  1.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.0,
		name.Anticipation: 1.5,
		name.Teamwork:     1.5,
		name.Defence:      0.8,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
	name.FR: {
		name.Tackle: 1.0, // technical factors
		name.Mark:   1.0,
		name.Positioning:  1.0, // mental factors
		name.Decision:     1.5,
		name.Workrate:     2.0,
		name.Anticipation: 1.5,
		name.Teamwork:     1.5,
		name.Defence:      0.8,
		name.Speed:    1.0, // physical factors
		name.Strength: 1.5,
	},
}

