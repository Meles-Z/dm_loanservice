package utils

import (
	"fmt"
	"unicode"
)

// password policies
func Password(password string) (bool, error) {
	if len(password) < 10 {
		return false, fmt.Errorf("password must be at least 10 characters long")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return false, fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return false, fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return false, fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return false, fmt.Errorf("password must contain at least one special character")
	}

	return true, nil
}
