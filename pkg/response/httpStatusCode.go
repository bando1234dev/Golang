package response

const (
	ErrCodeSuccess = 20001 // Success
	ErrParamInvalid = 20003 // Email is invalid
)

var msg = map[int]string {
	ErrCodeSuccess: "success",
	ErrParamInvalid: "Email is invalid",
}