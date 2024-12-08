package utils

import (
	"errors"
	"regexp"
)

// ValidateIndianMobileNumber validates an Indian mobile number
func ValidateIndianMobileNumber(mobile string) error {
	// Regular expression for valid Indian mobile numbers
	regex := `^[6-9]\d{9}$`

	matched, err := regexp.MatchString(regex, mobile)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("invalid Indian mobile number")
	}

	return nil
}
