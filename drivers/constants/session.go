package constants

type state string

var (
	ForgotPasswordCurrent state = "current"
	ForgotPasswordNext    state = "next"
)

type Flow map[state]string

func (f Flow) IsValid(curr, next string) bool {
	if f == nil {
		return false
	}
	return f[ForgotPasswordCurrent] == curr && f[ForgotPasswordNext] == next
}

const (
	ForgotPasswordInit           = "init"
	ForgotPasswordValidateCode   = "validateCode"
	ForgotPasswordValidateMobile = "validateMobile"
	ForgotPasswordValidateOTP    = "validateOTP"
	ForgotPasswordSetPassword    = "setPassword"
)

var (
	ForgotPasswordFlow = map[string]Flow{
		ForgotPasswordValidateCode: {
			ForgotPasswordCurrent: ForgotPasswordInit,
			ForgotPasswordNext:    ForgotPasswordValidateCode,
		},
		ForgotPasswordValidateMobile: {
			ForgotPasswordCurrent: ForgotPasswordValidateMobile,
			ForgotPasswordNext:    ForgotPasswordValidateOTP,
		},
		ForgotPasswordValidateOTP: {
			ForgotPasswordCurrent: ForgotPasswordValidateMobile,
			ForgotPasswordNext:    ForgotPasswordValidateOTP,
		},
		ForgotPasswordSetPassword: {
			ForgotPasswordCurrent: ForgotPasswordValidateOTP,
			ForgotPasswordNext:    ForgotPasswordSetPassword,
		},
	}

	ForgotPasswordState = map[string]Flow{
		ForgotPasswordInit: {
			ForgotPasswordCurrent: ForgotPasswordInit,
			ForgotPasswordNext:    ForgotPasswordValidateCode,
		},
		ForgotPasswordValidateMobile: {
			ForgotPasswordCurrent: ForgotPasswordValidateMobile,
			ForgotPasswordNext:    ForgotPasswordValidateOTP,
		},
		ForgotPasswordValidateOTP: {
			ForgotPasswordCurrent: ForgotPasswordValidateOTP,
			ForgotPasswordNext:    ForgotPasswordSetPassword,
		},
	}
)
