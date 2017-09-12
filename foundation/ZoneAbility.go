package foundation

import "sort"

const (
	AB_SPEED = "Speed"
	AB_HEALTH = "Health"
	AB_WORKRATE = "Workrate"

	AB_DRIBBLE = "Dribble"
	AB_PASS = "Pass"
	AB_CROSS = "CrossPass"
	AB_SHOOT = "Shoot"
	AB_LONGSHOOT = "LongShoot"

	AB_ANTICIPATION = "Anticipation"
	AB_POSITIONING = "Positioning"
	AB_OFFTHEBALL = "Offtheball"
	AB_TEAMWORK = "Teamwork"
	AB_DECISION = "Decision"
	AB_DEFENCE = "Defence"
	AB_ATTACK = "Attack"

)

type ZoneAbility struct {
	Name       string
	Attributes map[string]*Attribute
}

func (this ZoneAbility) Avg() int {
	return this.Sum() / len(this.Attributes)
}

var PossessionZoneFactors = map[string]PossessionAttrFactor{
	MC: PossessionAttrFactor{
		Factor: 2.0,
		AttrFactors: map[string]float32 {
			// technical factors
			AB_PASS: 3.0,
			AB_DRIBBLE: 1.0,
			// mental factors
			AB_OFFTHEBALL: 1.5,
			AB_DECISION: 1.2,
			AB_WORKRATE: 2.0,
			AB_ANTICIPATION: 1.2,
			// physical factors
			AB_SPEED: 1.2,
		},
	},
}
type PossessionAttrFactor struct {
	Factor      float32
	AttrFactors map[string]float32
}
/*
var PossesionZoneFactor = map[string]PossesionAttributeFactor {
	MC: PossesionAttributeFactor{
		// technical factors
		AB_PASS: 3.0,
		AB_DRIBBLE: 1.0,
		// mental factors
		AB_OFFTHEBALL: 1.5,
		AB_DECISION: 1.2,
		AB_WORKRATE: 2.0,
		AB_ANTICIPATION: 1.2,
		// physical factors
		AB_SPEED: 1.2,
	},

}*/
func (this ZoneAbility) Posession() int {
	var sum float32
	for key, factor := range PossessionZoneFactors[this.Name].AttrFactors {
		attirubte, ok := this.Attributes[key]
		if ok {
			sum += float32(attirubte.Point) * factor
		}
	}
	sum *= PossessionZoneFactors[this.Name].Factor
	return int(sum)
}
func (this ZoneAbility) Defence() {

}
func (this ZoneAbility) Offence() {

}
func (this ZoneAbility) Sum() int {
	sum := 0
	for _, v := range this.Attributes {
		sum += v.Point
	}
	return sum
}

func (this ZoneAbility) Sort() []*Attribute {
	array := []*Attribute{}
	for _, el := range this.Attributes {
		array = append(array, el)
	}
	sort.Sort(sortable(array))
	return array
}
func (this ZoneAbility) Effect(debuff ZoneAbility) {

}
func (this ZoneAbility) Clone() ZoneAbility {
	clone := ZoneAbility{ Name: this.Name, Attributes: map[string]*Attribute{} }
	for k, v := range this.Attributes {
		ability := *v
		clone.Attributes[k] = &ability
	}
	return clone
}
func (this ZoneAbility) StrongerThan(other ZoneAbility) bool {
	return this.Sum() > other.Sum()
}
