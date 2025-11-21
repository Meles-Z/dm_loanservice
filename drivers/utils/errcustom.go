package utils

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound   = errors.New("resource not found")
	ErrBadRequest = errors.New("bad request")
)

var (
	ErrorInvalidRequest = &Error{Code: "91", Message: "invalid request"}
	ErrorGeneral        = &Error{Code: "99", Message: "general service error"}

	ErrorPermissionDenied = &Error{Code: "401", Message: "permission denied"}
)

var (
	ErrorNotFound          = NewError("90", "resource not found")
	ErrorNewInvalidRequest = NewError("91", "invalid request")
	ErrorNewDatabase       = NewError("92", "error database")
	ErrorNewDataNotFound   = NewError("93", "data not found")
	ErrorNewParsingData    = NewError("94", "failed parsing data")
	ErrorNewGeneral        = NewError("99", "general service error")
	ErrorActionNotAllowed  = NewError("98", "action not allowed")

	ErrorNewLogin    = NewError("4000", "invalid username, password")
	ErrorInvalidRole = NewError("4004", "invalid role")

	ErrorNewUserNotFound        = NewError("4001", "User not found")
	ErrorNewEmailExist          = NewError("4020", "email already exist")
	ErrorNewCategoryDetailExist = NewError("4030", "category name already exist")
	ErrorNewCategoryFieldExist  = NewError("4040", "category field name already exist")
	ErrorNewRoleExist           = NewError("4050", "role not exist")
	ErrorNewRoleNotFound        = NewError("4051", "role not found")
	ErrorNewLimitReqResetPass   = NewError("4100", "request to reset password has reached its limit")
	ErrorNewChangePassword      = NewError("4101", "failed to change password")
	ErrorNewOtpNotFound         = NewError("4200", "Otp not found")
	ErrorNewMaxOtpNotMatch      = NewError("4201", "You have reached the maximum number of attempts. Please wait for 10 minutes")
	ErrorNewMaxSendOTP          = NewError("4203", "You have reached the maximum number of attempts. Please wait for 20 minute")
	ErrorNewCreateOtp           = NewError("4204", "failed create otp")
	ErrorNewFailedSendSms       = NewError("4205", "failed send sms")
	ErrorInvalidOldPassword     = NewError("4206", "invalid old password")
	ErrorInvalidNewPassword     = NewError("4207", "same password cant be the same with old password")

	ErrorFailedDecodeFile = NewError("102", "failed decode file upload")
	ErrorInvalidFileType  = NewError("103", "invalid file type")

	ErrorNewToken         = NewError("2000", "failed token")
	ErrorNewGenerateToken = NewError("2001", "failed generate token")
	ErrorNewTokenExpired  = NewError("2002", "token expired")
	ErrNewMaxOtpNotMatch  = NewError("2003", "max otp not match")
	ErrNewOtpExpired      = NewError("2004", "otp has Expired")
	ErrNewOtpOtpNotMatch  = NewError("2005", "otp not match")

	ErrorNewSendEmail  = NewError("3000", "failed send email")
	ErrorTeamNotFound  = NewError("4051", "team not found")
	ErrorCodeExpired   = NewError("4300", "code has expired")
	ErrorStateInvalid  = NewError("4301", "invalid state")
	ErrorMobileInvalid = NewError("4302", "invalid code/mobile number")

	ErrorMFAAlreadyEnabled  = NewError("5000", "MFA is already enabled")
	ErrorMFAAlreadyDisabled = NewError("5001", "MFA is already disabled")
	ErrorMFAEnabled         = NewError("5002", "MFA has been enabled")
	ErrorMFADisabled        = NewError("5003", "MFA has been disabled")

	ErrorMFANotEnabled    = NewError("5004", "Enable your MFA first")
	ErrorMFARequired      = NewError("200", "MFA required, please verify")
	ErrorMFASetupRequired = NewError("200", "MFA setup required, please setup MFA")
)

type Error struct {
	Service string      `json:"service"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Error   error       `json:"-"`
	Detail  interface{} `json:"detail"`
}

func (e Error) String() string {
	// causing nil pointer return fmt.Sprintf("%s:%s:%s", e.Service, e.Message, e.Error.Error())
	return fmt.Sprintf("%s:%s", e.Service, e.Message)
}

func NewError(errorCode string, message string) error {
	return &ApplicationError{
		Code:    errorCode,
		Message: message,
	}
}

type ApplicationError struct {
	Code    string
	Message string
}

func (e *ApplicationError) Error() string {
	return e.Message
}
