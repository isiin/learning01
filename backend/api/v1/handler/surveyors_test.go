package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_QuerySurveyors(t *testing.T) {
	// Ginをテストモードに設定（ログ出力を抑制）
	gin.SetMode(gin.TestMode)

	// テストケースの定義
	tests := []struct {
		name       string
		queryParam string // テストするクエリパラメータ (?q=...)
		wantStatus int    // 期待するHTTPステータスコード
	}{
		{
			name:       "正常系: パラメータあり",
			queryParam: "?q=001",
			wantStatus: http.StatusOK,
		},
		{
			name:       "異常系: パラメータなし",
			queryParam: "",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "異常系: パラメータが空",
			queryParam: "?q=",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 1. レスポンスを記録するレコーダーとコンテキストを作成
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 2. リクエストを擬似的に作成してセット
			c.Request, _ = http.NewRequest("GET", "/surveyors/query"+tt.queryParam, nil)

			// 3. ハンドラーを実行
			h := NewHandler()
			h.QuerySurveyors(c)

			// 4. 結果の検証
			if w.Code != tt.wantStatus {
				t.Errorf("status code got %d, want %d", w.Code, tt.wantStatus)
			}

			// 必要であればレスポンスボディ(w.Body.String())の検証もここで行えます
		})
	}
}
