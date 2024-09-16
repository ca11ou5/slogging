package slogging

import (
	"crypto/rand"
	"encoding/hex"
)

func generateTraceId() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
