package core

import "sort"

const (
	AB_SPEED = "Speed"
	AB_HEALTH = "Health"
	AB_WORKRATE = "Workrate"
	AB_STRENGTH = "Strength"

	AB_DRIBBLE = "Dribble"
	AB_PASS = "Pass"
	AB_CROSS = "CrossPass"
	AB_SHOOT = "Shoot"
	AB_LONGSHOOT = "LongShoot"

	AB_TACKLE = "Tackle"
	AB_MARK = "Mark"
	AB_ANTICIPATION = "Anticipation"
	AB_POSITIONING = "Positioning"
	AB_OFFTHEBALL = "Offtheball"
	AB_TEAMWORK = "Teamwork"
	AB_DECISION = "Decision"
	AB_DEFENCE = "Defence"
	AB_ATTACK = "Attack"

)

type ZoneAbility struct {
	Name       Zone
	Attributes map[string]*Attribute
}

func (this ZoneAbility) Avg() int {
	return this.Sum() / len(this.Attributes)
}


func (this ZoneAbility) calculateAttr(key string, factor float32) float32 {
	attribute, ok := this.Attributes[key]
	if ok {
		return float32(attribute.Point) * factor
	} else {
		return 0
	}
}
func (this ZoneAbility) Scoring() int {
	var sum float32
	for key, factor := range ScoringFactor.AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	return int(sum)
}
func (this ZoneAbility) Posession() int {
	return this.calculate(PossessionFactors)
}
func (this ZoneAbility) Defence() int {
	return this.calculate(DefenceFactors)
}
func (this ZoneAbility) Offence() int {
	return this.calculate(OffenceFactors)
}
func (this ZoneAbility) calculate(positionAttrFactors ZoneAttrFactors) int {
	var sum float32
	for key, factor := range positionAttrFactors[this.Name].AttrFactors {
		sum += this.calculateAttr(key, factor)
	}
	sum *= positionAttrFactors[this.Name].Factor
	return int(sum)
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
