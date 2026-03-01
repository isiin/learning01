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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetSurveyors_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		oid      string
		mockRet  domain.Surveyors
		expected []GetSurveyorsResponse
	}{
		{
			name:     "Empty",
			oid:      "aa",
			mockRet:  domain.Surveyors(nil),
			expected: []GetSurveyorsResponse{},
		},
		{
			name: "Success",
			oid:  "bb",
			mockRet: domain.Surveyors{
				{ID: "11111", Name: "サンプル1"},
				{ID: "22222", Name: "サンプル2"},
			},
			expected: []GetSurveyorsResponse{
				{ID: "11111", Name: "サンプル1"},
				{ID: "22222", Name: "サンプル2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("GET", "/dummy?office-id="+tt.oid, nil)

			uc := new(MockSurveyUseCase)
			uc.On("GetSurveyors",
				// mockに渡されるパラメータの検証はここに書く
				mock.MatchedBy(func(filter domain.SurveyorFilter) bool {
					return filter.OfficeID == tt.oid
				}),
			).Return(tt.mockRet, nil)

			GetSurveyors(uc)(c)

			assert := assert.New(t)

			assert.Equal(http.StatusOK, w.Code)
			assert.Empty(c.Errors)

			expectedJson, _ := json.Marshal(tt.expected)
			assert.Equal(string(expectedJson), w.Body.String())

		})
	}
}

func Test_GetSurveyors_Validation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		q  string // テストするパラメータ (?q=...)
		ok bool   // 想定結果 true:検証成功、false:検証エラー
	}{
		{q: "", ok: true},
		{q: "?office-id=", ok: true},
		{q: "?office-id=X1", ok: true},
		{q: "?office-id=123", ok: false},
		{q: "?office-id=あ", ok: false},
	}

	for _, tt := range tests {
		t.Run("param:"+tt.q, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			req, _ := http.NewRequest("GET", "/dummy"+tt.q, nil)
			c.Request = req

			uc := new(MockSurveyUseCase)
			if tt.ok {
				// 失敗ケースでモックを設定しないことで「バリデーションエラー時はUseCaseが呼ばれないこと」も暗黙的に検証できる
				uc.On("GetSurveyors", mock.Anything).
					Return(domain.Surveyors{}, nil)
			}

			GetSurveyors(uc)(c)

			assert := assert.New(t)

			if tt.ok {
				assert.Equal(http.StatusOK, w.Code)
				assert.Empty(c.Errors)
			} else {
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

func Test_GetSurveyors_FailureLogic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/dummy?office-id=XX", nil)

	// ドメインロジックがエラーを返す想定
	uc := new(MockSurveyUseCase)
	uc.On("GetSurveyors", mock.Anything).
		Return(domain.Surveyors(nil), errs.NewBusinessError(errs.Exclusion))

	GetSurveyors(uc)(c)

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
type MockSurveyUseCase struct {
	mock.Mock
}

func (m *MockSurveyUseCase) GetSurveyors(filter domain.SurveyorFilter) (domain.Surveyors, error) {
	args := m.Called(filter)
	return args.Get(0).(domain.Surveyors), args.Error(1)
}
