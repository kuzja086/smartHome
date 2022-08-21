package apperror

import "encoding/json"

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
	Code             string `json:"code"`
}

const (
	InternalError   = "US-000"
	notConfirmPassc = "US-001"
	emptyUsername   = "US-002"
	emptyPassword   = "US-003"
	HashGen         = "US-004"
	userNotFound    = "US-005"
	userExists      = "US-006"
	ErrorCreateUser = "US-007"
	authFaild       = "US-008"
)

var (
	NotConfirmPass = NewAppError("confirm password is wrong", "", notConfirmPassc, nil)
	EmptyUsername  = NewAppError("emppty username", "", emptyUsername, nil)
	EmptyPassword  = NewAppError("empty password", "", emptyPassword, nil)
	UserNotFound   = NewAppError("user by username not fond", "", userNotFound, nil)
	UserExists     = NewAppError("user already exists", "", userExists, nil)
	AuthFaild      = NewAppError("Auth is faild", "", authFaild, nil)
)

func NewAppError(message, developerMessage, code string, err error) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func systemError(err error) *AppError {
	return NewAppError("internal system error", "API error", InternalError, err)
}
