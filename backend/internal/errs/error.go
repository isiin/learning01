package errs

import "fmt"

// 業務エラーを表現するエラー
type BusinessError struct {
	code    ErrorCode
	details []string
}

func NewBusinessError(errCode ErrorCode, details ...string) *BusinessError {
	return &BusinessError{
		code:    errCode,
		details: details,
	}
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s details: %v", e.code, e.code.GetMessage(), e.details)
}

func (e *BusinessError) GetCode() ErrorCode {
	return e.code
}

func (e *BusinessError) GetDetails() []string {
	return e.details
}

// 想定外のエラーを表現するエラー
type SystemError struct {
	message string
	cause   error
}

func NewSystemError(message string, cause error) *SystemError {
	return &SystemError{
		message: message,
		cause:   cause,
	}
}

func (e *SystemError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%s] %s cause: %v", Internal, e.message, e.cause)
	}
	return fmt.Sprintf("[%s] %s", Internal, e.message)
}

func (e *SystemError) Unwrap() error {
	return e.cause
}
