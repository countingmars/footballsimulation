package misc

import "encoding/json"

func Json(object interface{}) string {
	out, _ := json.Marshal(object)
	return string(out)
}
