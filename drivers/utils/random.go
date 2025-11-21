package utils

import (
	crp "crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func GenerateThreadId() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	uniqueID := ulid.MustNew(ulid.Timestamp(t), entropy)
	return uniqueID.String()
}

// Checks if digits are sequential ascending/descending
func isSequential(code string) bool {
	for i := 1; i < len(code); i++ {
		prev := int(code[i-1] - '0')
		curr := int(code[i] - '0')
		if curr != prev+1 && curr != prev-1 {
			return false
		}
	}
	return true
}

// Checks if digits form an arithmetic pattern (like 1357, 8642)
func isArithmeticPattern(code string) bool {
	if len(code) < 2 {
		return false
	}
	diff := int(code[1]-'0') - int(code[0]-'0')
	for i := 1; i < len(code); i++ {
		currDiff := int(code[i]-'0') - int(code[i-1]-'0')
		if currDiff != diff {
			return false
		}
	}
	return true
}

// Generate an ultra-secure 8-digit passcode
func GenerateUltraSecurePasscode() (string, error) {
	for {
		digits := []rune("0123456789")
		passcode := make([]rune, 0, 8)

		for i := 0; i < 10; i++ {
			nBig, err := crp.Int(crp.Reader, big.NewInt(int64(len(digits))))
			if err != nil {
				return "", err
			}
			idx := nBig.Int64()
			passcode = append(passcode, digits[idx])
			// Remove used digit to avoid duplicates
			digits = append(digits[:idx], digits[idx+1:]...)
		}

		code := string(passcode)
		if !isSequential(code) && !isArithmeticPattern(code) {
			return code, nil
		}
		// retry if pattern detected
	}
}

func GenerateMFASecret() (string, error) {
	const secretSize = 20 // 160 bits
	buf := make([]byte, secretSize)

	// Use crypto/rand.Reader for cryptographically secure randomness.
	if _, err := io.ReadFull(crp.Reader, buf); err != nil {
		return "", fmt.Errorf("failed to generate secret: %w", err)
	}

	enc := base32.StdEncoding.WithPadding(base32.NoPadding)
	return enc.EncodeToString(buf), nil
}
