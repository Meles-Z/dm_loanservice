package utils

import "context"

func IsAuthorized(ctx context.Context, request interface{}) bool {
	// TODO: Dial to Authorization Service
	// TODO: Handle error
	return true
}
