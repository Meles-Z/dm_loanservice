package utils

import (
	"time"

	OTP "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Otp interface {
	GenerateCodeCustom(secret string, t time.Time) (string, error)
	ValidateCustom(passcode, secret string, t time.Time) (bool, error)
}

func NewOtp(period uint) *opts {
	return &opts{opt: totp.ValidateOpts{
		Period:    period * 60,
		Digits:    OTP.DigitsSix,
		Algorithm: OTP.AlgorithmSHA1,
	}}
}

type opts struct {
	opt totp.ValidateOpts
}

// GenerateCodeCustom extends github.com/pquerna/otp/totp GenerateCodeCustom
func (o *opts) GenerateCodeCustom(secret string, t time.Time) (string, error) {
	return totp.GenerateCodeCustom(secret, t, o.opt)
}

// ValidateCustom extends github.com/pquerna/otp/totp ValidateCustom
func (o *opts) ValidateCustom(passcode, secret string, t time.Time) (bool, error) {
	return totp.ValidateCustom(passcode, secret, t, o.opt)
}
