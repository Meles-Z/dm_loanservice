package constants

const (
	CsrfID = "gorilla.csrf.Token"

	MAXOTP                   int32 = 2
	MaxNotMatchOtp           int32 = 3
	MaxRequestForgotPassword       = 100
	MaxOTPLockRequestTime          = 20
	MaxOTPLockErrorTime            = 10

	SuccessCode = "00"
	Success     = "Success"

	BucketThProfiles  = "th-profiles"
	SubscriptionBasic = "BASIC"
	UserEnable        = "ENABLED"
	UserDisable       = "DISABLED"
	UserRequested     = "REQUESTED"
)

const (
	RoleAdmin       = 1
	RoleOps         = 7
	RoleUserRegular = 2
	RoleUserPremium = 3
	RolePublic      = 5
)

// allowed as admin or ops
var AllowedAccessMap = map[int64]struct{}{
	RoleAdmin: {},
	RoleOps:   {},
}

const (
	// action user update status
	ActionUserUpdateStatusEnabled  = 1
	ActionUserUpdateStatusDisabled = 2
)

var AllowedUserAction = map[int64]map[int]string{
	RoleAdmin: {
		ActionUserUpdateStatusEnabled:  UserEnable,
		ActionUserUpdateStatusDisabled: UserDisable,
	},
}
