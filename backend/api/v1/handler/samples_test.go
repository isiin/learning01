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
	tests := []struct {
		q string // テストするパラメータ (?q=...)
		s int    // 期待するHTTPステータスコード
	}{
		{q: "?q1=001&q2=002&q3=" + uuid.New().String() + "&q4=003-004-005&q4=a@gmail.com&q5=1,2,3&q6=2023-01-01", s: http.StatusOK},
		{q: "?q1=001&q2=02", s: http.StatusOK},
		{q: " ", s: http.StatusBadRequest},
		{q: "?q1=", s: http.StatusBadRequest},
		{q: "?q1=12", s: http.StatusBadRequest},
		{q: "?q1=001&q2=0021", s: http.StatusBadRequest},
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
