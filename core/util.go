package core

import "encoding/json"

func Json(data map[string]interface{}) string {
	result, _ := json.Marshal(data)
	return string(result)
}
