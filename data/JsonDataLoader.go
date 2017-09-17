package data

import (
	"encoding/json"
	"reflect"
	"github.com/countingmars/fb/base/player"
)

func LoadPlayer(filename string) *player.Player {
	var player player.Player
	var bytes = LoadFile(filename)
	json.Unmarshal(bytes, &player)
	return &player
}
func LoadJsonFile(filename string, destin interface{}) {
	if reflect.ValueOf(destin).Kind() != reflect.Ptr {
		panic("destin must be a ptr")
	}
	bytes := LoadFile(filename)
	json.Unmarshal(bytes, destin)
}