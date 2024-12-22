package respone

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Invalid
	ErrInvalidToken     = 30001 // InvalidToken
	// REGISTER CODE
	ErrCodeUserHasExist = 5001 //  user has already exist

)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrCodeUserHasExist: "User has already exist",
}
