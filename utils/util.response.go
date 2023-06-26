package utils

import "encoding/json"

func ErrorMessage(err string) string {
	msg := map[string]string {
		"Message": err,
	}

	jsonStr, _ := json.Marshal(msg)
	return string(jsonStr)
}