package apperror

import "encoding/json"

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
	Code             string `json:"code"`
}

const (
	internalError   = "US-000"
	notConfirmPassc = "US-001"
	emptyUsername   = "US-002"
	HashGen         = "US-003"
)

var (
	NotConfirmPass = NewAppError("confirm password is wrong", "", notConfirmPassc, nil)
	EmptyUsername  = NewAppError("emppty username", "", emptyUsername, nil)
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
	return NewAppError("internal system error", "API error", internalError, err)
}
