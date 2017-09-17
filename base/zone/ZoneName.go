package zone

import (
	"github.com/countingmars/fb/base"
)

var Names = struct {
	DR, DL, DC,
	MR, ML, MC,
	FR, FL, FC,
	GK base.Name
}{
	GK: base.Name("GK"),
	DR: base.Name("DR"),
	DL: base.Name("DL"),
	DC: base.Name("DC"),
	MR: base.Name("MR"),
	ML: base.Name("ML"),
	MC: base.Name("MC"),
	FR: base.Name("FR"),
	FL: base.Name("FL"),
	FC: base.Name("FC"),
}
