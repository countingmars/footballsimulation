package base


type ZoneAttrFactor struct {
	Factor      float32
	AttrFactors map[string]float32
}
type ZoneAttrFactors map[Zone]ZoneAttrFactor

var DefenceFactors = ZoneAttrFactors{
	Zones.DC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 2.0,
			AB_MARK: 3.0,

			// mental factors
			AB_POSITIONING: 3.0,
			AB_DECISION: 2.0,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 3.0,
			AB_DEFENCE: 3.0,
			// physical factors
			AB_SPEED: 1.2,
			AB_STRENGTH: 3.0,
		},
	},
	Zones.DL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 2.5,
			AB_MARK: 3.0,

			// mental factors
			AB_POSITIONING: 2.5,
			AB_DECISION: 2.0,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 2.0,
			AB_DEFENCE: 2.0,
			// physical factors
			AB_SPEED: 2.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.DR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 2.5,
			AB_MARK: 3.0,

			// mental factors
			AB_POSITIONING: 2.5,
			AB_DECISION: 2.0,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 2.0,
			AB_DEFENCE: 2.0,
			// physical factors
			AB_SPEED: 2.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.MC: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.5,
			AB_MARK: 2.0,

			// mental factors
			AB_POSITIONING: 2.5,
			AB_DECISION: 2.0,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 2.5,
			AB_DEFENCE: 1.5,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.ML: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.0,
			AB_MARK: 1.5,

			// mental factors
			AB_POSITIONING: 1.5,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 1.5,
			AB_DEFENCE: 1.0,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.MR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.0,
			AB_MARK: 1.5,

			// mental factors
			AB_POSITIONING: 1.5,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 2.0,
			AB_TEAMWORK: 1.5,
			AB_DEFENCE: 1.0,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.0,
			AB_MARK: 1.0,

			// mental factors
			AB_POSITIONING: 1.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			AB_TEAMWORK: 1.5,
			AB_DEFENCE: 0.8,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.FL: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.0,
			AB_MARK: 1.0,

			// mental factors
			AB_POSITIONING: 1.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			AB_TEAMWORK: 1.5,
			AB_DEFENCE: 0.8,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
	Zones.FR: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_TACKLE: 1.0,
			AB_MARK: 1.0,

			// mental factors
			AB_POSITIONING: 1.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			AB_TEAMWORK: 1.5,
			AB_DEFENCE: 0.8,
			// physical factors
			AB_SPEED: 1.0,
			AB_STRENGTH: 1.5,
		},
	},
}
var OffenceFactors = ZoneAttrFactors{
	Zones.DC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 1.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.2,
			AB_ATTACK: 1.5,
			AB_TEAMWORK: 2.0,
		},
	},
	Zones.DL: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			AB_CROSS: 3.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 2.0,
			AB_ATTACK: 2.0,
			AB_TEAMWORK: 2.0,
		},
	},
	Zones.DR: ZoneAttrFactor{
		Factor: 1.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			AB_CROSS: 3.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 2.0,
			AB_ATTACK: 2.0,
			AB_TEAMWORK: 2.0,
		},
	},
	Zones.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 3.0,
			AB_DRIBBLE: 2.0,
			AB_CROSS: 1.5,
			AB_LONGSHOOT: 1.5,
			// mental factors
			AB_OFFTHEBALL: 2.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.5,
			AB_ANTICIPATION: 2.0,
			// physical factors
			AB_SPEED: 1.0,
			AB_ATTACK: 2.0,
			AB_TEAMWORK: 3.0,
		},
	},
	Zones.ML: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 3.0,
			AB_CROSS: 2.5,
			AB_LONGSHOOT: 1.2,
			// mental factors
			AB_OFFTHEBALL: 2.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.5,
			AB_ANTICIPATION: 2.0,
			// physical factors
			AB_SPEED: 3.0,
			AB_ATTACK: 2.0,
			AB_TEAMWORK: 1.5,
		},
	},
	Zones.MR: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 3.0,
			AB_CROSS: 2.5,
			AB_LONGSHOOT: 1.2,
			// mental factors
			AB_OFFTHEBALL: 2.0,
			AB_DECISION: 1.5,
			AB_WORKRATE: 2.5,
			AB_ANTICIPATION: 2.0,
			// physical factors
			AB_SPEED: 3.0,
			AB_ATTACK: 2.0,
			AB_TEAMWORK: 1.5,
		},
	},
	Zones.FC: ZoneAttrFactor{
		Factor: 3.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 3.0,
			AB_CROSS: 2.0,
			AB_LONGSHOOT: 1.0,
			// mental factors
			AB_OFFTHEBALL: 3.0,
			AB_DECISION: 3.0,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 2.5,
			// physical factors
			AB_SPEED: 2.0,
			AB_ATTACK: 3.0,
			AB_TEAMWORK: 1.5,
		},
	},
	Zones.FL: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 3.0,
			AB_CROSS: 2.5,
			AB_LONGSHOOT: 1.0,
			// mental factors
			AB_OFFTHEBALL: 3.0,
			AB_DECISION: 2.5,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 2.0,
			// physical factors
			AB_SPEED: 2.0,
			AB_ATTACK: 3.0,
			AB_TEAMWORK: 1.5,
		},
	},
	Zones.FR: ZoneAttrFactor{
		Factor: 2.5,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 3.0,
			AB_CROSS: 2.5,
			AB_LONGSHOOT: 1.0,
			// mental factors
			AB_OFFTHEBALL: 3.0,
			AB_DECISION: 2.5,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 2.0,
			// physical factors
			AB_SPEED: 2.0,
			AB_ATTACK: 3.0,
			AB_TEAMWORK: 1.5,
		},
	},
}
var ScoringFactor = ZoneAttrFactor{
	Factor: 1.0,
	AttrFactors: map[string]float32 {
		// technical factors
		AB_DRIBBLE: 1.2,
		AB_SHOOT: 3.0,
		// mental factors
		AB_OFFTHEBALL: 1.5,
		AB_DECISION: 1.5,
		AB_ANTICIPATION: 1.5,
		AB_ATTACK: 1.5,
		// physical factors
	},
}


var PossessionFactors = map[Zone]ZoneAttrFactor{
	Zones.DC: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 2.0,
			AB_DRIBBLE: 1.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.2,
		},
	},
	Zones.DL: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.5,
		},
	},
	Zones.DR: ZoneAttrFactor{
		Factor: 1.2,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.5,
		},
	},
	Zones.MC: ZoneAttrFactor{
		Factor: 2.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 3.0,
			AB_DRIBBLE: 1.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.2,
		},
	},
	Zones.ML: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 2.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 2.0,
		},
	},
	Zones.MR: ZoneAttrFactor{
		Factor: 1.3,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 2.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 2.0,
		},
	},
	Zones.FC: ZoneAttrFactor{
		Factor: 1.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.5,
		},
	},
	Zones.FL: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.5,
		},
	},
	Zones.FR: ZoneAttrFactor{
		Factor: 0.8,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 1.5,
			AB_DRIBBLE: 1.5,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 1.5,
			AB_ANTICIPATION: 1.5,
			// physical factors
			AB_SPEED: 1.5,
		},
	},
}


