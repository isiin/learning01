package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestHandler_GetSamples(t *testing.T) {
	// Ginをテストモードに設定（ログ出力を抑制）
	gin.SetMode(gin.TestMode)

	// テストケースの定義
	uuid := uuid.New().String()
	email := "a@example.com"

	tests := []struct {
		q string // テストするパラメータ (?q=...)
		s int    // 期待するHTTPステータスコード
	}{
		// Q1       string    `form:"q1" binding:"required,alphanum,len=3"`
		{q: "?q1=123", s: http.StatusOK},
		{q: "?q1=12", s: http.StatusBadRequest},
		{q: "?q1=1234", s: http.StatusBadRequest},
		{q: "?q1=あ", s: http.StatusBadRequest},
		{q: "?q1=", s: http.StatusBadRequest},
		{q: " ", s: http.StatusBadRequest},

		// Q2       string    `form:"q2" binding:"omitempty,max=3,min=2"`
		{q: "?q1=001&q2=123", s: http.StatusOK},
		{q: "?q1=001&q2=12", s: http.StatusOK},
		{q: "?q1=001&q2=あい", s: http.StatusOK},
		{q: "?q1=001&q2=", s: http.StatusOK},
		{q: "?q1=001&q2=1", s: http.StatusBadRequest},
		{q: "?q1=001&q2=1234", s: http.StatusBadRequest},

		// UUID     string    `form:"uuid" binding:"omitempty,uuid"`
		{q: "?q1=001&uuid=" + uuid, s: http.StatusOK},
		{q: "?q1=001&uuid=1234", s: http.StatusBadRequest},

		// Email    string    `form:"email" binding:"omitempty,email"`
		{q: "?q1=001&email=" + email, s: http.StatusOK},
		{q: "?q1=001&email=a@example", s: http.StatusBadRequest},

		// IntArray []int     `form:"intArray" collection_format:"csv"`
		{q: "?q1=001&intArray=1,2,3", s: http.StatusOK},
		{q: "?q1=001&intArray=1", s: http.StatusOK},
		{q: "?q1=001&intArray=a,b", s: http.StatusBadRequest},

		// DateUtc  time.Time `form:"dateUtc" time_format:"2006-01-02" time_utc:"1"`
		{q: "?q1=001&dateUtc=2023-01-01", s: http.StatusOK},
		{q: "?q1=001&dateUtc=2023-02-31", s: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.q, func(t *testing.T) {
			// 1. レスポンスを記録するレコーダーとコンテキストを作成
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 2. リクエストを擬似的に作成してセット
			url := "/samples" + tt.q
			c.Request, _ = http.NewRequest("GET", url, nil)

			// 3. ハンドラーを実行
			h := NewHandler()
			h.GetSamples(c)

			// 4. 結果の検証
			if w.Code != tt.s {
				t.Errorf("status code got %d, want %d", w.Code, tt.s)
			}

			// 必要であればレスポンスボディ(w.Body.String())の検証もここで行えます
			log.Println(w.Body.String())
		})
	}
}
