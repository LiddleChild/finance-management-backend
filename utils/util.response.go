package utils

import "encoding/json"

func JSONMessage(str string) string {
	msg := map[string]string {
		"Message": str,
	}

	jsonStr, _ := json.Marshal(msg)
	return string(jsonStr)
}