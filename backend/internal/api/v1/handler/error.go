package handler

import (
	"errors"
	"fmt"
	"log"
	"react-ts/backend/internal/errs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse エラーレスポンスの構造体
// Swaggerのドキュメント生成用に各ハンドラーから参照されることを想定しています
type ErrorResponse struct {
	Code    string   `json:"code" example:"INVALID_REQUEST"`
	Message string   `json:"message" example:"不正なリクエストです"`
	Details []string `json:"details" example:"aaa,bbb,ccc"`
}

// ErrorHandler はgin.ErrorTypePublicであるエラーが検出された場合クライアントにJSONエラーレスポンスを送信して処理を中断するmiddleware。
// 各ハンドラーはエラー処理で「c.Error(err).SetType(gin.ErrorTypePublic)」を呼ぶ必要がある
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			cd := errs.Internal
			details := []string{}
			var e *errs.BusinessError
			if errors.As(err, &e) {
				cd = e.GetCode()
				details = e.GetDetails()
			} else {
				// 業務エラーでない場合はログを出力する
				log.Printf("System Error: %v\n", err.Err)
			}

			res := ErrorResponse{
				Code:    string(cd),
				Message: cd.GetMessage(),
				Details: details,
			}
			c.AbortWithStatusJSON(cd.GetStatus(), res)
		}

	}
}

// TODO createValidationDetails はバリデーションエラーから詳細なメッセージのスライスを生成します。
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
