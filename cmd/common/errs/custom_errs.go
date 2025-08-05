package errs

const (
	ErrBadRequest    = 400_101
	ErrMsgBadRequest = "Bad Request"

	ErrUnauthorized    = 401_101
	ErrMsgUnauthorized = "Unauthorized"

	ErrInvalidToken    = 401_102
	ErrMsgInvalidToken = "Invalid Token"
	ErrForbidden       = 403_101
	ErrMsgForbidden    = "Forbidden"

	ErrNotFound          = 404_101
	ErrMsgNotFound       = "Not Found"
	ErrInternalServer    = 500_101
	ErrMsgInternalServer = "Internal Server Error"

	// Authentication specific errors
	ErrUsernameExists    = 400_102
	ErrMsgUsernameExists = "Username already exists"

	ErrInvalidCredentials    = 401_103
	ErrMsgInvalidCredentials = "Invalid credentials"

	ErrTokenAlreadyRevoked    = 400_103
	ErrMsgTokenAlreadyRevoked = "Token already revoked"

	// Upload specific errors
	ErrFileTooLarge    = 400_104
	ErrMsgFileTooLarge = "File size exceeds 8MB limit"

	ErrInvalidFileType    = 400_105
	ErrMsgInvalidFileType = "Only image files are allowed"

	ErrFileProcessing    = 500_102
	ErrMsgFileProcessing = "Failed to process uploaded file"
)

var ErrService = map[int]string{
	ErrBadRequest:          ErrMsgBadRequest,
	ErrForbidden:           ErrMsgForbidden,
	ErrNotFound:            ErrMsgNotFound,
	ErrInternalServer:      ErrMsgInternalServer,
	ErrUnauthorized:        ErrMsgUnauthorized,
	ErrInvalidToken:        ErrMsgInvalidToken,
	ErrUsernameExists:      ErrMsgUsernameExists,
	ErrInvalidCredentials:  ErrMsgInvalidCredentials,
	ErrTokenAlreadyRevoked: ErrMsgTokenAlreadyRevoked,
	ErrFileTooLarge:        ErrMsgFileTooLarge,
	ErrInvalidFileType:     ErrMsgInvalidFileType,
	ErrFileProcessing:      ErrMsgFileProcessing,
}

type CustomError struct {
	Code       int    `json:"code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"err_msg"`
}

func (c *CustomError) Error() string {
	return c.Message
}

// NewCustomErrWithMsg receive 2 args
func NewCustomErrWithMsg(code int, errMsg string) *CustomError {
	return &CustomError{
		Code:       code,
		StatusCode: code / 1000,
		Message:    errMsg,
	}
}

// NewCustomError return instance which construct from code, and message was got from defined constant
// if code doesn't define, it returns default value
func NewCustomError(code int) *CustomError {
	msg, ok := ErrService[code]
	if !ok {
		msg = ErrMsgInternalServer
		code = ErrInternalServer
	}
	return &CustomError{
		Code:       code,
		StatusCode: code / 1000,
		Message:    msg,
	}
}
