package zone

import "github.com/countingmars/fb/base/attr"

type ZoneAttrFactor struct {
	Factor      float32
	AttrFactors map[attr.Name]float32
}
type ZoneAttrFactors map[Name]ZoneAttrFactor

var DefenceFactors = ZoneAttrFactors{
	Names.DC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 2.0,
			attr.Names.Mark: 3.0,

			// mental factors
			attr.Names.Positioning: 3.0,
			attr.Names.Decision: 2.0,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 3.0,
			attr.Names.Defence: 3.0,
			// physical factors
			attr.Names.Speed: 1.2,
			attr.Names.Strength: 3.0,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 2.5,
			attr.Names.Mark: 3.0,

			// mental factors
			attr.Names.Positioning: 2.5,
			attr.Names.Decision: 2.0,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 2.0,
			attr.Names.Defence: 2.0,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 2.5,
			attr.Names.Mark: 3.0,

			// mental factors
			attr.Names.Positioning: 2.5,
			attr.Names.Decision: 2.0,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 2.0,
			attr.Names.Defence: 2.0,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.5,
			attr.Names.Mark: 2.0,

			// mental factors
			attr.Names.Positioning: 2.5,
			attr.Names.Decision: 2.0,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 2.5,
			attr.Names.Defence: 1.5,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.0,
			attr.Names.Mark: 1.5,

			// mental factors
			attr.Names.Positioning: 1.5,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 1.5,
			attr.Names.Defence: 1.0,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.0,
			attr.Names.Mark: 1.5,

			// mental factors
			attr.Names.Positioning: 1.5,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 2.0,
			attr.Names.Teamwork: 1.5,
			attr.Names.Defence: 1.0,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.0,
			attr.Names.Mark: 1.0,

			// mental factors
			attr.Names.Positioning: 1.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			attr.Names.Teamwork: 1.5,
			attr.Names.Defence: 0.8,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.0,
			attr.Names.Mark: 1.0,

			// mental factors
			attr.Names.Positioning: 1.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			attr.Names.Teamwork: 1.5,
			attr.Names.Defence: 0.8,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Tackle: 1.0,
			attr.Names.Mark: 1.0,

			// mental factors
			attr.Names.Positioning: 1.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			attr.Names.Teamwork: 1.5,
			attr.Names.Defence: 0.8,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Strength: 1.5,
		},
	},
}
var OffenceFactors = ZoneAttrFactors{
	Names.DC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 1.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.2,
			attr.Names.Attack: 1.5,
			attr.Names.Teamwork: 2.0,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			attr.Names.Cross: 3.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Attack: 2.0,
			attr.Names.Teamwork: 2.0,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			attr.Names.Cross: 3.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Attack: 2.0,
			attr.Names.Teamwork: 2.0,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 3.0,
			attr.Names.Dribble: 2.0,
			attr.Names.Cross: 1.5,
			attr.Names.LongShoot: 1.5,
			// mental factors
			attr.Names.OffTheBall: 2.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.5,
			attr.Names.Anticipation: 2.0,
			// physical factors
			attr.Names.Speed: 1.0,
			attr.Names.Attack: 2.0,
			attr.Names.Teamwork: 3.0,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 3.0,
			attr.Names.Cross: 2.5,
			attr.Names.LongShoot: 1.2,
			// mental factors
			attr.Names.OffTheBall: 2.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.5,
			attr.Names.Anticipation: 2.0,
			// physical factors
			attr.Names.Speed: 3.0,
			attr.Names.Attack: 2.0,
			attr.Names.Teamwork: 1.5,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 3.0,
			attr.Names.Cross: 2.5,
			attr.Names.LongShoot: 1.2,
			// mental factors
			attr.Names.OffTheBall: 2.0,
			attr.Names.Decision: 1.5,
			attr.Names.Workrate: 2.5,
			attr.Names.Anticipation: 2.0,
			// physical factors
			attr.Names.Speed: 3.0,
			attr.Names.Attack: 2.0,
			attr.Names.Teamwork: 1.5,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 3.0,
			attr.Names.Cross: 2.0,
			attr.Names.LongShoot: 1.0,
			// mental factors
			attr.Names.OffTheBall: 3.0,
			attr.Names.Decision: 3.0,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 2.5,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Attack: 3.0,
			attr.Names.Teamwork: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 3.0,
			attr.Names.Cross: 2.5,
			attr.Names.LongShoot: 1.0,
			// mental factors
			attr.Names.OffTheBall: 3.0,
			attr.Names.Decision: 2.5,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 2.0,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Attack: 3.0,
			attr.Names.Teamwork: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 3.0,
			attr.Names.Cross: 2.5,
			attr.Names.LongShoot: 1.0,
			// mental factors
			attr.Names.OffTheBall: 3.0,
			attr.Names.Decision: 2.5,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 2.0,
			// physical factors
			attr.Names.Speed: 2.0,
			attr.Names.Attack: 3.0,
			attr.Names.Teamwork: 1.5,
		},
	},
}
var ScoringFactor = ZoneAttrFactor{
	Factor: 1.0,
	AttrFactors: map[attr.Name]float32 {
		// technical factors
		attr.Names.Dribble: 1.2,
		attr.Names.Shoot: 3.0,
		// mental factors
		attr.Names.OffTheBall: 1.5,
		attr.Names.Decision: 1.5,
		attr.Names.Anticipation: 1.5,
		attr.Names.Attack: 1.5,
		// physical factors
	},
}


var PossessionFactors = map[Name]ZoneAttrFactor{
	Names.DC: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 2.0,
			attr.Names.Dribble: 1.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.2,
		},
	},
	Names.DL: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.5,
		},
	},
	Names.DR: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.5,
		},
	},
	Names.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 3.0,
			attr.Names.Dribble: 1.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.2,
		},
	},
	Names.ML: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 2.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 2.0,
		},
	},
	Names.MR: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 2.0,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 2.0,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 2.0,
		},
	},
	Names.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.5,
		},
	},
	Names.FL: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.5,
		},
	},
	Names.FR: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[attr.Name]float32 {
			// technical factors
			attr.Names.Pass: 1.5,
			attr.Names.Dribble: 1.5,
			// mental factors
			attr.Names.OffTheBall: 1.5,
			attr.Names.Decision: 1.2,
			attr.Names.Workrate: 1.5,
			attr.Names.Anticipation: 1.5,
			// physical factors
			attr.Names.Speed: 1.5,
		},
	},
}


