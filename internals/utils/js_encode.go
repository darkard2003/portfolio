package utils

import "encoding/json"

func JsEncode(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "'encoding error'"
	}
	return string(b)
}
