package zone

import "strings"

type Name string

func (this Name) In(line string) bool {
	return strings.Contains(string(this), line)
}

var Names = struct{
	DR, DL, DC,
	MR, ML, MC,
	FR, FL, FC,
	GK Name
} {
	GK: Name("GK"),
	DR: Name("DR"),
	DL: Name("DL"),
	DC: Name("DC"),
	MR: Name("MR"),
	ML: Name("ML"),
	MC: Name("MC"),
	FR: Name("FR"),
	FL: Name("FL"),
	FC: Name("FC"),
}