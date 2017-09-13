package gamesimulator

var GoalScenarios = [...]Highlight {
	{ "GOAL", "SUPER PLAY" },
	{ "GOAL", "PK" },
	{ "GOAL", "FK" },
	{ "GOAL", "CK" },
	{ "GOAL", "TEAM PLAY" },
	{ "GOAL", "LONG SHOOT" },
	{ "GOAL", "LUCKY" },
	{ "GOAL", "MISS" },
	{ "GOAL", "OWN" },
	{ "GOAL", "COUNTER" },
}

type Highlight struct {
	Type string
	Message string
}

