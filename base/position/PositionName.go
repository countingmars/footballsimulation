package position

import (
	"github.com/countingmars/fb/base"
)
var Names = struct {
	GK,
	DR, DL, DC,
	WBL, WBR, DMC,
	ML, MR, MC,
	AML, AMR, AMC,
	STC base.Name
} {
	GK: base.Name("GK"),
	DR: base.Name("DR"), DL: base.Name("DL"), DC: base.Name("DC"),
	WBL: base.Name("WBL"), WBR: base.Name("WBR"), DMC: base.Name("DMC"),
	ML: base.Name("ML"), MR: base.Name("MR"), MC: base.Name("MC"),
	AML: base.Name("AML"), AMR: base.Name("AMR"), AMC: base.Name("AMC"),
	STC: base.Name("STC"),
}