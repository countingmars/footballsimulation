package base

import "strings"

type Zone string

func (this Zone) In(line string) bool {
	return strings.Contains(string(this), line)
}

type zones struct{
	GK         Zone
	DR, DL, DC Zone
	MR, ML, MC Zone
	FR, FL, FC Zone
}
var Zones = zones{
	GK: "GK",
	DR: "DR",
	DL: "DL",
	DC: "DC",
	MR: "MR",
	ML: "ML",
	MC: "MC",
	FR: "FR",
	FL: "FL",
	FC: "FC",
}