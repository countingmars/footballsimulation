package foundation


type PositionAttrFactor struct {
	Factor      float32
	AttrFactors map[string]float32
}
type PositionAttrFactors map[string]PositionAttrFactor

var DefenceFactors = PositionAttrFactors{
	DC: PositionAttrFactor{
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
	DL: PositionAttrFactor{
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
	DR: PositionAttrFactor{
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
	MC: PositionAttrFactor{
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
	ML: PositionAttrFactor{
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
	MR: PositionAttrFactor{
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
	FC: PositionAttrFactor{
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
	FL: PositionAttrFactor{
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
	FR: PositionAttrFactor{
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
var OffenceFactors = PositionAttrFactors{
	DC: PositionAttrFactor{
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
	DL: PositionAttrFactor{
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
	DR: PositionAttrFactor{
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
	MC: PositionAttrFactor{
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
	ML: PositionAttrFactor{
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
	MR: PositionAttrFactor{
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
	FC: PositionAttrFactor{
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
	FL: PositionAttrFactor{
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
	FR: PositionAttrFactor{
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
var ScoringFactor = PositionAttrFactor{
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


var PossessionFactors = map[string]PositionAttrFactor{
	DC: PositionAttrFactor{
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
	DL: PositionAttrFactor{
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
	DR: PositionAttrFactor{
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
	MC: PositionAttrFactor{
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
	ML: PositionAttrFactor{
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
	MR: PositionAttrFactor{
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
	FC: PositionAttrFactor{
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
	FL: PositionAttrFactor{
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
	FR: PositionAttrFactor{
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


