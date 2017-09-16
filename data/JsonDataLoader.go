package data

import (
	"encoding/json"
	"github.com/countingmars/fb/base"
	"github.com/countingmars/fb/misc"
	"reflect"

)

func LoadTeam(filename string) *base.Team {
	var team base.Team
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