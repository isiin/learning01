package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"react-ts/backend/internal/domain"
	"react-ts/backend/internal/errs"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetSamples_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		mockRet  domain.Samples
		expected []GetSampleResponse
	}{
		{
			name:     "Empty",
			mockRet:  domain.Samples{},
			expected: []GetSampleResponse{},
		},
		{
			name: "Success",
			mockRet: domain.Samples{
				{ID: "11111", Name: "サンプル1"},
				{ID: "22222", Name: "サンプル2"},
			},
			expected: []GetSampleResponse{
				{ID: "11111", Name: "サンプル1"},
				{ID: "22222", Name: "サンプル2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("GET", "/dummy?q1=abc", nil)

			uc := new(MockSampleUseCase)
			uc.On("GetSamples").Return(tt.mockRet, nil)

			GetSamples(uc)(c)

			assert := assert.New(t)
			assert.Equal(http.StatusOK, w.Code)
			assert.Empty(c.Errors)

			expectedJson, _ := json.Marshal(tt.expected)
			assert.JSONEq(string(expectedJson), w.Body.String())
		})
	}
}

func Test_GetSamples_Validation(t *testing.T) {
	// Ginをテストモードに設定（ログ出力を抑制）
	gin.SetMode(gin.TestMode)

	// テストケースの定義
	tests := []struct {
		q  string // テストするパラメータ (?q=...)
		ok bool   // 想定結果 true:検証成功、false:検証エラー
	}{
		// Q1       string    `form:"q1" binding:"required,alphanum,len=3"`
		{q: "?q1=123", ok: true},
		{q: "?q1=12", ok: false},
		{q: "?q1=1234", ok: false},
		{q: "?q1=あ", ok: false},
		{q: "?q1=", ok: false},
		{q: " ", ok: false},

		// Q2       string    `form:"q2" binding:"omitempty,max=3,min=2"`
		{q: "?q1=001&q2=123", ok: true},
		{q: "?q1=001&q2=12", ok: true},
		{q: "?q1=001&q2=あい", ok: true},
		{q: "?q1=001&q2=", ok: true},
		{q: "?q1=001&q2=1", ok: false},
		{q: "?q1=001&q2=1234", ok: false},

		// UUID     string    `form:"uuid" binding:"omitempty,uuid"`
		{q: "?q1=001&uuid=" + uuid.New().String(), ok: true},
		{q: "?q1=001&uuid=1234", ok: false},

		// Email    string    `form:"email" binding:"omitempty,email"`
		{q: "?q1=001&email=a@example.com", ok: true},
		{q: "?q1=001&email=a@example", ok: false},

		// IntArray []int     `form:"intArray" collection_format:"csv"`
		{q: "?q1=001&intArray=1,2,3", ok: true},
		{q: "?q1=001&intArray=1", ok: true},
		{q: "?q1=001&intArray=a,b", ok: false},

		// DateUtc  time.Time `form:"dateUtc" time_format:"2006-01-02" time_utc:"1"`
		{q: "?q1=001&dateUtc=2023-01-01", ok: true},
		{q: "?q1=001&dateUtc=2023-02-31", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.q, func(t *testing.T) {
			// 1. レスポンスを記録するレコーダーとコンテキストを作成
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 2. リクエストを擬似的に作成してセット
			c.Request, _ = http.NewRequest("GET", "/dummy"+tt.q, nil)

			// 3. モックの設定とハンドラー実行
			uc := new(MockSampleUseCase)
			uc.On("GetSamples").Return(domain.Samples{}, nil)

			GetSamples(uc)(c)

			// 4. 結果の検証
			assert := assert.New(t)

			if tt.ok {
				// 成功ケースでは、エラーが積まれていないことを検証する
				assert.Equal(http.StatusOK, w.Code)
				assert.Empty(c.Errors)
			} else {
				// エラーケースでは、エラーの内容を検証する。
				pe := c.Errors.ByType(gin.ErrorTypePublic).Last()
				if assert.NotEmpty(pe) {
					var b *errs.BusinessError
					if errors.As(pe.Err, &b) {
						assert.Equal(errs.InvalidRequest, pe.Err.(*errs.BusinessError).GetCode())
					} else {
						assert.Fail("エラーコードが想定外です")
					}
				}
			}
		})
	}
}

func Test_GetSamples_FailureLogic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/dummy?q1=abc", nil)

	// ドメインロジックがエラーを返す想定
	uc := new(MockSampleUseCase)
	uc.On("GetSamples").Return(domain.Samples{}, errs.NewBusinessError(errs.Exclusion))

	GetSamples(uc)(c)

	assert := assert.New(t)

	// エラーの内容を検証する。
	pe := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if assert.NotEmpty(pe) {
		var b *errs.BusinessError
		if errors.As(pe.Err, &b) {
			assert.Equal(errs.Exclusion, pe.Err.(*errs.BusinessError).GetCode())
		} else {
			assert.Fail("エラーコードが想定外です")
		}
	}
}

// testify/mockを使用してモック作成
type MockSampleUseCase struct {
	mock.Mock
}

func (m *MockSampleUseCase) GetSamples() (domain.Samples, error) {
	args := m.Called()
	return args.Get(0).(domain.Samples), args.Error(1)
}
