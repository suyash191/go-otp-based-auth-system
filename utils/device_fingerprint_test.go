package utils

import (
	"fmt"
	"testing"
)

func TestGenerateFingerprint(t *testing.T) {
	// call the function twice. Output should be same
	deviceDetails := "s22" + "samsung" + "x64"
	fingerprint1 := GenerateFingerprint(deviceDetails)
	fingerprint2 := GenerateFingerprint(deviceDetails)
	if fingerprint1 != fingerprint2 {
		t.Errorf("Expected both fingerprints to be same, but got %s and %s", fingerprint1, fingerprint2)
	}
	fmt.Println("Both the fingerprints are same: ", fingerprint1)
}
