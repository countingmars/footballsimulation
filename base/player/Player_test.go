package player

import (
	"testing"
	"encoding/json"
	"strings"
	"github.com/countingmars/fb/base/position"
)

func TestPlayer_json(t *testing.T) {
	player := TestPlayer(position.DC)

	out, _ := json.Marshal(&player)
	actual := string(out)

	if strings.Contains(actual, "Name") {
		t.Error("json field names should be lowercase, but ", actual)
	}
}

