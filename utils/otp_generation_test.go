package utils

import (
	"fmt"
	"testing"
)

func TestOTPGeneration(t *testing.T) {
	otp := GenerateOTP()
	fmt.Println("Generated OTP: ", otp)
	if len(otp) != 6 {
		t.Errorf("Expected OTP to be of length 6, but got %d", len(otp))
	}
}
