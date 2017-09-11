package foundation

import "sort"

const (
	AB_SPEED = "Speed"
	AB_HEALTH = "Health"

	AB_DRIBBLE = "Dribble"
	AB_PASS = "Pass"
	AB_CROSS = "CrossPass"
	AB_SHOOT = "Shoot"
	AB_LONGSHOOT = "LongShoot"

	AB_DEFENCE = "Defence"
	AB_ATTACK = "Attack"
	AB_STRATEGY = "Strategy"
)

type ZoneAbility map[string]*Ability

func (this ZoneAbility) Avg() int {
	return this.Sum() / len(this)
}

func (this ZoneAbility) Sum() int {
	sum := 0
	for _, v := range this {
		sum += v.Point
	}
	return sum
}

func (this ZoneAbility) Sort() []*Ability{
	array := []*Ability{}
	for _, el := range this {
		array = append(array, el)
	}
	sort.Sort(sortable(array))
	return array
}
func (this ZoneAbility) Debuff(debuff int) {
	if debuff == 0 {
		return
	}
	remains := debuff % len(this)
	point := debuff / len(this)
	for _, v := range this {
		if remains > 0 {
			v.Point -= 1
			remains--
		}
		v.Point -= point
	}
}
func (this ZoneAbility) Buff(buff int) {
	if buff == 0 {
		return
	}
	remains := buff % len(this)
	point := buff / len(this)
	for _, v := range this {
		if remains > 0 {
			v.Point += 1
			remains--
		}
		v.Point += point
	}
}
func (this ZoneAbility) Clone() ZoneAbility {
	clone := ZoneAbility{}
	for k, v := range this {
		ability := *v
		clone[k] = &ability
	}
	return clone
}
func (this ZoneAbility) StrongerThan(other ZoneAbility) bool {
	return this.Sum() > other.Sum()
}
