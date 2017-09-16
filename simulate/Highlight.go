package simulate



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
	Side Side
	Time int
	Type HighlightType
	Message string
}


type Highlights []Highlight

func (this Highlights) Goals(side Side) int {
	goals := 0
	for _, v := range this {
		if v.Side == side && v.Type == Goal {
			goals++
		}
	}
	return goals
}
