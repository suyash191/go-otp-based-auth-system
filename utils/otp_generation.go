package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOTP generates a 6-digit OTP
func GenerateOTP() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return fmt.Sprintf("%06d", r.Intn(1000000))
}
