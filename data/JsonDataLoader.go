package data

import (
	"encoding/json"
	. "github.com/countingmars/fb/foundation"
	. "github.com/countingmars/fb/misc"
	"reflect"
	"fmt"
)

func LoadTeam(filename string) *Team {
	var team Team
	var bytes = LoadFile(filename)
	json.Unmarshal(bytes, &team)
	return &team
}
func LoadJsonFile(filename string, destin interface{}) {
	if reflect.ValueOf(destin).Kind() != reflect.Ptr {
		panic("destin must be a ptr")
	}
	bytes := LoadFile(filename)
	json.Unmarshal(bytes, destin)

	fmt.Println(string(bytes))
	fmt.Println(destin)
}