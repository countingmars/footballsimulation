package zone

import (
	"github.com/countingmars/fb/base/stats"
	"github.com/countingmars/fb/base"
)

type ZoneAttrFactor struct {
	Factor      float32
	AttrFactors map[base.Name]float32
}
type ZoneAttrFactors map[base.Name]ZoneAttrFactor

var DefenceFactors = ZoneAttrFactors{
	Names.DC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 2.0,
			stats.Names.Mark:   3.0,

			// mental factors
			stats.Names.Positioning:  3.0,
			stats.Names.Decision:     2.0,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     3.0,
			stats.Names.Defence:      3.0,
			// physical factors
			stats.Names.Speed:    1.2,
			stats.Names.Strength: 3.0,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 2.5,
			stats.Names.Mark:   3.0,

			// mental factors
			stats.Names.Positioning:  2.5,
			stats.Names.Decision:     2.0,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     2.0,
			stats.Names.Defence:      2.0,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 2.5,
			stats.Names.Mark:   3.0,

			// mental factors
			stats.Names.Positioning:  2.5,
			stats.Names.Decision:     2.0,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     2.0,
			stats.Names.Defence:      2.0,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.5,
			stats.Names.Mark:   2.0,

			// mental factors
			stats.Names.Positioning:  2.5,
			stats.Names.Decision:     2.0,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     2.5,
			stats.Names.Defence:      1.5,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.0,
			stats.Names.Mark:   1.5,

			// mental factors
			stats.Names.Positioning:  1.5,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     1.5,
			stats.Names.Defence:      1.0,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.0,
			stats.Names.Mark:   1.5,

			// mental factors
			stats.Names.Positioning:  1.5,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 2.0,
			stats.Names.Teamwork:     1.5,
			stats.Names.Defence:      1.0,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.0,
			stats.Names.Mark:   1.0,

			// mental factors
			stats.Names.Positioning:  1.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			stats.Names.Teamwork:     1.5,
			stats.Names.Defence:      0.8,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.0,
			stats.Names.Mark:   1.0,

			// mental factors
			stats.Names.Positioning:  1.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			stats.Names.Teamwork:     1.5,
			stats.Names.Defence:      0.8,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Tackle: 1.0,
			stats.Names.Mark:   1.0,

			// mental factors
			stats.Names.Positioning:  1.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			stats.Names.Teamwork:     1.5,
			stats.Names.Defence:      0.8,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Strength: 1.5,
		},
	},
}
var OffenceFactors = ZoneAttrFactors{
	Names.DC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    2.0,
			stats.Names.Dribble: 1.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed:    1.2,
			stats.Names.Attack:   1.5,
			stats.Names.Teamwork: 2.0,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			stats.Names.Cross:   3.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Attack:   2.0,
			stats.Names.Teamwork: 2.0,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			stats.Names.Cross:   3.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Attack:   2.0,
			stats.Names.Teamwork: 2.0,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      3.0,
			stats.Names.Dribble:   2.0,
			stats.Names.Cross:     1.5,
			stats.Names.LongShoot: 1.5,
			// mental factors
			stats.Names.OffTheBall:   2.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.5,
			stats.Names.Anticipation: 2.0,
			// physical factors
			stats.Names.Speed:    1.0,
			stats.Names.Attack:   2.0,
			stats.Names.Teamwork: 3.0,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      2.0,
			stats.Names.Dribble:   3.0,
			stats.Names.Cross:     2.5,
			stats.Names.LongShoot: 1.2,
			// mental factors
			stats.Names.OffTheBall:   2.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.5,
			stats.Names.Anticipation: 2.0,
			// physical factors
			stats.Names.Speed:    3.0,
			stats.Names.Attack:   2.0,
			stats.Names.Teamwork: 1.5,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      2.0,
			stats.Names.Dribble:   3.0,
			stats.Names.Cross:     2.5,
			stats.Names.LongShoot: 1.2,
			// mental factors
			stats.Names.OffTheBall:   2.0,
			stats.Names.Decision:     1.5,
			stats.Names.Workrate:     2.5,
			stats.Names.Anticipation: 2.0,
			// physical factors
			stats.Names.Speed:    3.0,
			stats.Names.Attack:   2.0,
			stats.Names.Teamwork: 1.5,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      2.0,
			stats.Names.Dribble:   3.0,
			stats.Names.Cross:     2.0,
			stats.Names.LongShoot: 1.0,
			// mental factors
			stats.Names.OffTheBall:   3.0,
			stats.Names.Decision:     3.0,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 2.5,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Attack:   3.0,
			stats.Names.Teamwork: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      2.0,
			stats.Names.Dribble:   3.0,
			stats.Names.Cross:     2.5,
			stats.Names.LongShoot: 1.0,
			// mental factors
			stats.Names.OffTheBall:   3.0,
			stats.Names.Decision:     2.5,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 2.0,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Attack:   3.0,
			stats.Names.Teamwork: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:      2.0,
			stats.Names.Dribble:   3.0,
			stats.Names.Cross:     2.5,
			stats.Names.LongShoot: 1.0,
			// mental factors
			stats.Names.OffTheBall:   3.0,
			stats.Names.Decision:     2.5,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 2.0,
			// physical factors
			stats.Names.Speed:    2.0,
			stats.Names.Attack:   3.0,
			stats.Names.Teamwork: 1.5,
		},
	},
}
var ScoringFactor = ZoneAttrFactor{
	Factor: 1.0,
	AttrFactors: map[base.Name]float32 {
		// technical factors
		stats.Names.Dribble: 1.2,
		stats.Names.Shoot:   3.0,
		// mental factors
		stats.Names.OffTheBall:   1.5,
		stats.Names.Decision:     1.5,
		stats.Names.Anticipation: 1.5,
		stats.Names.Attack:       1.5,
		// physical factors
	},
}


var PossessionFactors = map[base.Name]ZoneAttrFactor{
	Names.DC: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    2.0,
			stats.Names.Dribble: 1.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.2,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.5,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.5,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    3.0,
			stats.Names.Dribble: 1.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.2,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 2.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 2.0,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 2.0,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     2.0,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 2.0,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[base.Name]float32 {
			// technical factors
			stats.Names.Pass:    1.5,
			stats.Names.Dribble: 1.5,
			// mental factors
			stats.Names.OffTheBall:   1.5,
			stats.Names.Decision:     1.2,
			stats.Names.Workrate:     1.5,
			stats.Names.Anticipation: 1.5,
			// physical factors
			stats.Names.Speed: 1.5,
		},
	},
}


