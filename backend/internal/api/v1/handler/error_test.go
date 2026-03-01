package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"react-ts/backend/internal/errs"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name             string
		err              error
		expectedStatus   int
		expectedResponse ErrorResponse
	}{
		{
			name:           "BusinessError",
			err:            errs.NewBusinessError(errs.Exclusion, "xxxx", "yyyy"),
			expectedStatus: errs.Exclusion.GetStatus(),
			expectedResponse: ErrorResponse{
				Code:    string(errs.Exclusion),
				Message: errs.Exclusion.GetMessage(),
				Details: []string{"xxxx", "yyyy"},
			},
		},
		{
			name:           "SystemError",
			err:            errs.NewSystemError("zzzzz", errors.New("error")),
			expectedStatus: 500,
			expectedResponse: ErrorResponse{
				Code:    string(errs.Internal),
				Message: errs.Internal.GetMessage(),
				Details: []string{},
			},
		},
		{
			name:           "error",
			err:            errors.New("error"),
			expectedStatus: 500,
			expectedResponse: ErrorResponse{
				Code:    string(errs.Internal),
				Message: errs.Internal.GetMessage(),
				Details: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := gin.New()
			r.Use(ErrorHandler())

			// テスト用のハンドラーを定義
			r.GET("/dummy", func(c *gin.Context) {
				c.Error(tt.err).SetType(gin.ErrorTypePublic)
			})

			// リクエストを作成して実行
			req, _ := http.NewRequest("GET", "/dummy", nil)
			r.ServeHTTP(w, req)

			assert := assert.New(t)

			assert.Equal(tt.expectedStatus, w.Code)
			expectedJson, _ := json.Marshal(tt.expectedResponse)
			assert.JSONEq(string(expectedJson), w.Body.String())
		})
	}
}
