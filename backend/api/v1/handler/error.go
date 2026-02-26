package handler

import (
	"errors"
	"fmt"
	"react-ts/backend/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TODO エラーレスポンスの形
type ErrorResponse struct {
	Code    string   `json:"code" example:"INVALID_REQUEST"`
	Message string   `json:"message" example:"不正なリクエストです"`
	Details []string `json:"details" example:"[aaa,bbb,ccc]"`
}

// TODO エラー情報を格納したJSONレスポンスを返す
func ResponseErrorJson(c *gin.Context, status int, errCode domain.ErrorCode, details []string) {
	er := ErrorResponse{
		Code:    string(errCode),
		Message: errCode.GetMessage(),
		Details: details,
	}
	c.JSON(status, er)
}

// createValidationDetails はバリデーションエラーから詳細なメッセージのスライスを生成します。
func createValidationDetails(err error) []string {
	var ve validator.ValidationErrors
	// エラーがバリデーションエラーではない場合は、そのままエラーメッセージを返す
	if !errors.As(err, &ve) {
		return []string{err.Error()}
	}

	details := make([]string, len(ve))
	for i, fe := range ve {
		// ここでタグに応じて日本語のメッセージを返すことも可能です
		details[i] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fe.Field(), fe.Tag())
	}

	return details
}
