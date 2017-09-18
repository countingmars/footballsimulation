package zone

import (
	"github.com/countingmars/fb/base/name"
	"github.com/countingmars/fb/simulater/zone/effect"
)

type ZoneFactor struct {
	Factor      float32
	AttrFactors map[name.Name]float32
}
func (this ZoneFactor) AttrFormulas(zone *Zone) []effect.AttrFormula {
	list := []effect.AttrFormula{}
	for attrName, attrFactor := range this.AttrFactors {
		list = append(list, effect.AttrFormula{zone.Attribute(attrName), attrFactor})
	}
	return list
}

type ZoneFactors map[name.Name]ZoneFactor

func (this ZoneFactors) Formula(zone *Zone) *effect.Formula {
	zf := this[zone.Name]
	formula := effect.Formula{zf.AttrFormulas(zone), zf.Factor}

	return &formula
}

var DefenceFactors = ZoneFactors{
	name.DC: ZoneFactor{
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
	name.DL: ZoneFactor{
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
	name.DR: ZoneFactor{
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
	name.MC: ZoneFactor{
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
	name.ML: ZoneFactor{
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
	name.MR: ZoneFactor{
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
	name.FC: ZoneFactor{
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
	name.FL: ZoneFactor{
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
	name.FR: ZoneFactor{
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
var OffenceFactors = ZoneFactors{
	name.DC: ZoneFactor{
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
	name.DL: ZoneFactor{
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
	name.DR: ZoneFactor{
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
	name.MC: ZoneFactor{
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
	name.ML: ZoneFactor{
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
	name.MR: ZoneFactor{
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
	name.FC: ZoneFactor{
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
	name.FL: ZoneFactor{
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
	name.FR: ZoneFactor{
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

var ScoringFactor = ZoneFactor{
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


var PossessionFactors = ZoneFactors{
	name.DC: ZoneFactor{
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
	name.DL: ZoneFactor{
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
	name.DR: ZoneFactor{
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
	name.MC: ZoneFactor{
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
	name.ML: ZoneFactor{
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
	name.MR: ZoneFactor{
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
	name.FC: ZoneFactor{
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
	name.FL: ZoneFactor{
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
	name.FR: ZoneFactor{
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


