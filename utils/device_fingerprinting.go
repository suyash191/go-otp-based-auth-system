package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GenerateFingerprint generates a unique fingerprint for the device
func GenerateFingerprint(deviceDetails string) string {
	hash := sha256.Sum256([]byte(deviceDetails))
	return hex.EncodeToString(hash[:])
}
