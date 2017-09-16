package attr

type Name string

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

	Defence, Attack, Gk Name
} {
	Speed: Name("speed"), Health: Name("health"),  Workrate: Name("workrate"), 
	Strength: Name("strength"), 

	Dribble: Name("dribble"),  Pass: Name("pass"),  Cross: Name("cross"),  Shoot: Name("shoot"), 
	LongShoot: Name("longshoot"),  Tackle: Name("tackle"), 
	Mark: Name("mark"),  Finish: Name("finish"),  Heading: Name("heading"), 

	Anticipation: Name("anticipation"),  Creative: Name("creative"),  Positioning: Name("positioning"), 
	OffTheBall: Name("offtheball"),  Teamwork: Name("teamwork"),  Decision: Name("decision"),
	Strategy: Name("strategy"),

	Defence: Name("defence"),  Attack: Name("attack"),  Gk: Name("gk"), 
}