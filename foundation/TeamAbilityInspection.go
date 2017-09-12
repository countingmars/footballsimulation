package foundation

const (
	TAI_BALANCE = "balance"
	TAI_ATTACK = "attack"
	TAI_PASS = "pass"
	TAI_DEFENCE = "defence"
	TAI_PRESS = "press"
	TAI_COUNTER = "counter"
)
type TeamAbilityInspection struct {
	Strength *Attribute
	Weakness *Attribute
	Style string
}

func (this TeamAbilityInspection) Diff() int {
	return this.Strength.Point - this.Weakness.Point
}