package player

import (
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/base/position"
	"github.com/countingmars/fb/base/stats"
	"encoding/json"
)

type Player struct {
	Name       string             `json:"name"`
	Age        int                `json:"age"`
	Country    base.Country       `json:"country"`
	Attributes stats.Attributes   `json:"attributes"`
	Positions  position.Positions `json:"positions"`
}

func (this *Player) Clone() *Player {
	return &Player{
		Name:       this.Name,
		Age:        this.Age,
		Country:    this.Country,
		Attributes: this.Attributes.Clone(),
		Positions:  this.Positions.Clone(),
	}
}
func (this *Player) Json() string {
	out, _ := json.Marshal(this)
	return string(out)
}
