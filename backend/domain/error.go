package domain

import "fmt"

// アプリ固有のエラーコード
type ErrorCode string

const (
	ECInvalidRequest ErrorCode = "INVALID_REQUEST"
	ECExclusion      ErrorCode = "EXCLUSION"
	ECInternal       ErrorCode = "INTERNAL"
)

// ユーザーに見せるメッセージ
var errorMessages = map[ErrorCode]string{
	ECInvalidRequest: "不正なリクエストです",
	ECExclusion:      "すでに削除されています",
	ECInternal:       "想定外のエラーが発生しました",
}

func (e ErrorCode) GetMessage() string {
	if msg, ok := errorMessages[e]; ok {
		return msg
	}
	return "不明なエラーです"
}

// 業務エラーを表現するエラー
type BusinessError struct {
	ErrorCode ErrorCode
	Details   []string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s details: %v", e.ErrorCode, e.ErrorCode.GetMessage(), e.Details)
}

func NewBusinessError(errCode ErrorCode, details ...string) *BusinessError {
	return &BusinessError{
		ErrorCode: errCode,
		Details:   details,
	}
}

// 想定外のエラーを表現するエラー
type SystemException struct {
	Message string
	Cause   error
}

func (e *SystemException) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s cause: %v", ECInternal, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", ECInternal, e.Message)
}

func (e *SystemException) Unwrap() error {
	return e.Cause
}

func NewSystemException(message string, cause error) *SystemException {
	return &SystemException{
		Message: message,
		Cause:   cause,
	}
}
