package simpleauth

import (
	"encoding/base64"
	"log"
)

func decodeString(str string) string {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Printf("failed to decode: \n %v", err)
		return ""
	}
	return string(b)
}

func encodeString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
