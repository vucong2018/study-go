package respone

const (
	ErrCodeSuccess = 20001 // Success
	ErrCodeParamInvalid = 20003 // Invalid
)

var msg = map[int]string{
	ErrCodeSuccess: "success",
	ErrCodeParamInvalid: "Email is invalid",
}