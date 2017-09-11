package gamesimulator

var Highlights = [...]Highlight {
	{ "SUPER PLAY" },
	{ "PK" },
	{ "FK" },
	{ "CK" },
	{ "TEAM PLAY" },
	{ "LONG SHOOT" },
	{ "LUCKY" },
	{ "MISS" },
	{ "OWN" },
	{ "COUNTER" },
}

type Highlight struct {
	Message string
}

