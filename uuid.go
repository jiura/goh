package goh

import (
	"crypto/rand"
	"fmt"
)

func Uuid_Generate() (string, error) {
	uuid := make([]byte, 16)

	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Set the version to 4 (random)
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	// Set the variant to RFC 4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 1

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
