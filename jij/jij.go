package jij

import (
	"encoding/json"
	"strings"
)

// Convert wants JSON string, returns Escaped JSON string
func Convert(text string) string {
	return strings.TrimSpace(
		strings.TrimPrefix(
			strings.TrimSuffix(
				jsonize(
					strings.TrimSpace(text),
				),
				"}",
			),
			"{\"text\":",
		),
	)
}

func jsonize(text string) string {
	dummy := struct {
		Text string `json:"text"`
	}{
		text,
	}
	val, err := json.Marshal(dummy)
	if err != nil {
		return ""
	}

	return string(val)
}
