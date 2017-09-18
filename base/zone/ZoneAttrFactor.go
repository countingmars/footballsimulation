package zone

import (
	"github.com/countingmars/fb/base/name"
)

type ZoneAttrFactor struct {
	Factor      float32
	AttrFactors map[name.Name]float32
}
type ZoneAttrFactors map[name.Name]ZoneAttrFactor

var DefenceFactors = ZoneAttrFactors{
	name.DC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 2.0,
			name.Mark:   3.0,

			// mental factors
			name.Positioning:  3.0,
			name.Decision:     2.0,
			name.Workrate:     1.5,
			name.Anticipation: 2.0,
			name.Teamwork:     3.0,
			name.Defence:      3.0,
			// physical factors
			name.Speed:    1.2,
			name.Strength: 3.0,
		},
	},
	name.DL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 2.5,
			name.Mark:   3.0,

			// mental factors
			name.Positioning:  2.5,
			name.Decision:     2.0,
			name.Workrate:     2.0,
			name.Anticipation: 2.0,
			name.Teamwork:     2.0,
			name.Defence:      2.0,
			// physical factors
			name.Speed:    2.0,
			name.Strength: 1.5,
		},
	},
	name.DR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 2.5,
			name.Mark:   3.0,

			// mental factors
			name.Positioning:  2.5,
			name.Decision:     2.0,
			name.Workrate:     2.0,
			name.Anticipation: 2.0,
			name.Teamwork:     2.0,
			name.Defence:      2.0,
			// physical factors
			name.Speed:    2.0,
			name.Strength: 1.5,
		},
	},
	name.MC: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.5,
			name.Mark:   2.0,

			// mental factors
			name.Positioning:  2.5,
			name.Decision:     2.0,
			name.Workrate:     2.0,
			name.Anticipation: 2.0,
			name.Teamwork:     2.5,
			name.Defence:      1.5,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
	name.ML: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.0,
			name.Mark:   1.5,

			// mental factors
			name.Positioning:  1.5,
			name.Decision:     1.5,
			name.Workrate:     2.0,
			name.Anticipation: 2.0,
			name.Teamwork:     1.5,
			name.Defence:      1.0,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
	name.MR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.0,
			name.Mark:   1.5,

			// mental factors
			name.Positioning:  1.5,
			name.Decision:     1.5,
			name.Workrate:     2.0,
			name.Anticipation: 2.0,
			name.Teamwork:     1.5,
			name.Defence:      1.0,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
	name.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.0,
			name.Mark:   1.0,

			// mental factors
			name.Positioning:  1.0,
			name.Decision:     1.5,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			name.Teamwork:     1.5,
			name.Defence:      0.8,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
	name.FL: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.0,
			name.Mark:   1.0,

			// mental factors
			name.Positioning:  1.0,
			name.Decision:     1.5,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			name.Teamwork:     1.5,
			name.Defence:      0.8,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
	name.FR: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Tackle: 1.0,
			name.Mark:   1.0,

			// mental factors
			name.Positioning:  1.0,
			name.Decision:     1.5,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			name.Teamwork:     1.5,
			name.Defence:      0.8,
			// physical factors
			name.Speed:    1.0,
			name.Strength: 1.5,
		},
	},
}
var OffenceFactors = ZoneAttrFactors{
	name.DC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    2.0,
			name.Dribble: 1.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed:    1.2,
			name.Attack:   1.5,
			name.Teamwork: 2.0,
		},
	},
	name.DL: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			name.Cross:   3.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed:    2.0,
			name.Attack:   2.0,
			name.Teamwork: 2.0,
		},
	},
	name.DR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			name.Cross:   3.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed:    2.0,
			name.Attack:   2.0,
			name.Teamwork: 2.0,
		},
	},
	name.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      3.0,
			name.Dribble:   2.0,
			name.Cross:     1.5,
			name.LongShoot: 1.5,
			// mental factors
			name.OffTheBall:   2.0,
			name.Decision:     1.5,
			name.Workrate:     2.5,
			name.Anticipation: 2.0,
			// physical factors
			name.Speed:    1.0,
			name.Attack:   2.0,
			name.Teamwork: 3.0,
		},
	},
	name.ML: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      2.0,
			name.Dribble:   3.0,
			name.Cross:     2.5,
			name.LongShoot: 1.2,
			// mental factors
			name.OffTheBall:   2.0,
			name.Decision:     1.5,
			name.Workrate:     2.5,
			name.Anticipation: 2.0,
			// physical factors
			name.Speed:    3.0,
			name.Attack:   2.0,
			name.Teamwork: 1.5,
		},
	},
	name.MR: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      2.0,
			name.Dribble:   3.0,
			name.Cross:     2.5,
			name.LongShoot: 1.2,
			// mental factors
			name.OffTheBall:   2.0,
			name.Decision:     1.5,
			name.Workrate:     2.5,
			name.Anticipation: 2.0,
			// physical factors
			name.Speed:    3.0,
			name.Attack:   2.0,
			name.Teamwork: 1.5,
		},
	},
	name.FC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      2.0,
			name.Dribble:   3.0,
			name.Cross:     2.0,
			name.LongShoot: 1.0,
			// mental factors
			name.OffTheBall:   3.0,
			name.Decision:     3.0,
			name.Workrate:     1.5,
			name.Anticipation: 2.5,
			// physical factors
			name.Speed:    2.0,
			name.Attack:   3.0,
			name.Teamwork: 1.5,
		},
	},
	name.FL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      2.0,
			name.Dribble:   3.0,
			name.Cross:     2.5,
			name.LongShoot: 1.0,
			// mental factors
			name.OffTheBall:   3.0,
			name.Decision:     2.5,
			name.Workrate:     1.5,
			name.Anticipation: 2.0,
			// physical factors
			name.Speed:    2.0,
			name.Attack:   3.0,
			name.Teamwork: 1.5,
		},
	},
	name.FR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:      2.0,
			name.Dribble:   3.0,
			name.Cross:     2.5,
			name.LongShoot: 1.0,
			// mental factors
			name.OffTheBall:   3.0,
			name.Decision:     2.5,
			name.Workrate:     1.5,
			name.Anticipation: 2.0,
			// physical factors
			name.Speed:    2.0,
			name.Attack:   3.0,
			name.Teamwork: 1.5,
		},
	},
}
var ScoringFactor = ZoneAttrFactor{
	Factor: 1.0,
	AttrFactors: map[name.Name]float32 {
		// technical factors
		name.Dribble: 1.2,
		name.Shoot:   3.0,
		// mental factors
		name.OffTheBall:   1.5,
		name.Decision:     1.5,
		name.Anticipation: 1.5,
		name.Attack:       1.5,
		// physical factors
	},
}


var PossessionFactors = map[name.Name]ZoneAttrFactor{
	name.DC: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    2.0,
			name.Dribble: 1.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.2,
		},
	},
	name.DL: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.5,
		},
	},
	name.DR: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.5,
		},
	},
	name.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    3.0,
			name.Dribble: 1.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.2,
		},
	},
	name.ML: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 2.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 2.0,
		},
	},
	name.MR: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 2.0,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     2.0,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 2.0,
		},
	},
	name.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.5,
		},
	},
	name.FL: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.5,
		},
	},
	name.FR: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[name.Name]float32 {
			// technical factors
			name.Pass:    1.5,
			name.Dribble: 1.5,
			// mental factors
			name.OffTheBall:   1.5,
			name.Decision:     1.2,
			name.Workrate:     1.5,
			name.Anticipation: 1.5,
			// physical factors
			name.Speed: 1.5,
		},
	},
}


