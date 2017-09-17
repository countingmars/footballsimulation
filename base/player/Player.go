package player

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/stats"
)

type Player struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Country base.Country `json:"country"`
	Ability stats.Ability `json:"ability"`
	Positions position.Positions `json:"positions"`
}

func (this *Player) Clone() *Player {
	return &Player{
		Name: this.Name,
		Age: this.Age,
		Country: this.Country,
		Ability: this.Ability.Clone(),
		Positions: this.Positions.Clone(),
	}
}