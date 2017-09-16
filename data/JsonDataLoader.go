package data

import (
	"encoding/json"
	"github.com/countingmars/fb/base/team"
	"github.com/countingmars/fb/misc"
	"reflect"
)

func LoadTeam(filename string) *team.Team {
	var team team.Team
	var bytes = misc.LoadFile(filename)
	json.Unmarshal(bytes, &team)
	return &team
}
func LoadJsonFile(filename string, destin interface{}) {
	if reflect.ValueOf(destin).Kind() != reflect.Ptr {
		panic("destin must be a ptr")
	}
	bytes := misc.LoadFile(filename)
	json.Unmarshal(bytes, destin)
}