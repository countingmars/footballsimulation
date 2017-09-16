package simulater



type HighlightType string

const (
	Goal HighlightType = "goal"
	Shoot HighlightType = "shoot"
)


var GoalHighlights = []Highlight {
	{ Type: Goal, Message: "SUPER PLAY" },
	{ Type: Goal, Message: "PK" },
	{ Type: Goal, Message: "FK" },
	{ Type: Goal, Message: "HEADING" },
	{ Type: Goal, Message: "TEAM PLAY" },
	{ Type: Goal, Message: "LONG SHOOT" },
	{ Type: Goal, Message: "LUCKY" },
	{ Type: Goal, Message: "MISS" },
	{ Type: Goal, Message: "OWN" },
	{ Type: Goal, Message: "COUNTER" },
}
var ShootHighlights = []Highlight {
	{ Type: Shoot, Message: "NEARLY GOAL" },
	{ Type: Shoot, Message: "IT'S A ROCKET TO MT.FUJI" },
	{ Type: Shoot, Message: "NICE FREEKICK AND CATCH" },
	{ Type: Shoot, Message: "NEVER MAKE A GOAL" },
	{ Type: Shoot, Message: "WHAT THE HELL OF THE SHOOT" },
	{ Type: Shoot, Message: "THAT'S FINISH ... LOOKS LIKE BAD CONDITION" },
	{ Type: Shoot, Message: "HE DOESN'T WANT A GOAL" },
	{ Type: Shoot, Message: "LOOKS LIKE A SPY" },
	{ Type: Shoot, Message: "TERRIBLE FINISH" },
}

type Highlight struct {
	Time       int
	Possession int
	Side       Side
	Type       HighlightType
	Message    string
}


type Highlights []Highlight

func (this Highlights) Possession(side Side) int {
	sum := 0
	for _, hl := range this {
		if hl.Side == side {
			sum += hl.Possession
		}
	}
	return sum
}

func (this Highlights) Goals(side Side) int {
	return this.sum(Goal, side)
}
func (this Highlights) Shoots(side Side) int {
	return this.sum(Shoot, side)
}
func (this Highlights) sum(hlType HighlightType, side Side) int {
	sum := 0
	for _, hl := range this {
		if hl.Side == side && hl.Type == hlType {
			sum++
		}
	}
	return sum
}
