package stats

import (
	"github.com/countingmars/fb/base"
)
var Names = struct {
	// physical
	Speed, Health, Workrate,
	Strength,

	// technical
	Dribble, Pass, Cross, Shoot,
	LongShoot, Tackle,
	Mark, Finish, Heading,

	// mental
	Anticipation, Creative, Positioning,
	OffTheBall, Teamwork, Decision,
	Strategy,

	Defence, Attack, Gk base.Name
} {
	Speed: base.Name("speed"), Health: base.Name("health"),  Workrate: base.Name("workrate"), 
	Strength: base.Name("strength"), 

	Dribble: base.Name("dribble"),  Pass: base.Name("pass"),  Cross: base.Name("cross"),  Shoot: base.Name("shoot"), 
	LongShoot: base.Name("longshoot"),  Tackle: base.Name("tackle"), 
	Mark: base.Name("mark"),  Finish: base.Name("finish"),  Heading: base.Name("heading"), 

	Anticipation: base.Name("anticipation"),  Creative: base.Name("creative"),  Positioning: base.Name("positioning"), 
	OffTheBall: base.Name("offtheball"),  Teamwork: base.Name("teamwork"),  Decision: base.Name("decision"),
	Strategy: base.Name("strategy"),

	Defence: base.Name("defence"),  Attack: base.Name("attack"),  Gk: base.Name("gk"), 
}