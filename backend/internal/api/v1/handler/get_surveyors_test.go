package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GetSurveyors_Success(t *testing.T) {
	// Ginをテストモードに設定（ログ出力を抑制）
	gin.SetMode(gin.TestMode)

	// TODO テストケースの定義
	tests := []struct {
		q string // テストするクエリパラメータ (?q=...)
		s int    // 期待するHTTPステータスコード
	}{
		{q: " ", s: 200},
		{q: "?id=000001&office-id=XX&view=detail", s: 200},
		{q: "?id=000001", s: 200},
		{q: "?office-id=XX", s: 200},
		{q: "?view=basic", s: 200},
		{q: "?view=detail", s: 200},
	}

	for _, tt := range tests {
		t.Run(tt.q, func(t *testing.T) {
			// 1. レスポンスを記録するレコーダーとコンテキストを作成
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 2. リクエストを擬似的に作成してセット
			c.Request, _ = http.NewRequest("GET", "/dummy"+tt.q, nil)

			// 3. ハンドラーを実行
			GetSurveyors()(c)

			// 4. 結果の検証
			if w.Code != tt.s {
				t.Errorf("status code got %d, want %d", w.Code, tt.s)
			}

			// 必要であればレスポンスボディ(w.Body.String())の検証もここで行えます
			// TODO モック作成とレスポンスボディの検証...
		})
	}
}

func TestGetSurveyors_Validation(t *testing.T) {
	//TODO
	// // Ginをテストモードに設定（ログ出力を抑制）
	// gin.SetMode(gin.TestMode)

	// // TODO テストケースの定義
	// tests := []struct {
	// 	q string // テストするクエリパラメータ (?q=...)
	// 	e error  // 期待するエラー
	// }{
	// 	{
	// 		q:    "?q=001", e:    http.StatusOK,
	// 	},
	// 	{
	// 		q:    "",
	// 		e:    http.StatusBadRequest,
	// 	},
	// 	{
	// 		q:    "?q=",
	// 		e:    http.StatusBadRequest,
	// 	},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		// 1. レスポンスを記録するレコーダーとコンテキストを作成
	// 		w := httptest.NewRecorder()
	// 		c, _ := gin.CreateTestContext(w)

	// 		// 2. リクエストを擬似的に作成してセット
	// 		c.Request, _ = http.NewRequest("GET", "/dummy"+tt.q, nil)

	// 		// 3. ハンドラーを実行
	// 		GetSurveyors()(c)

	// 		// 4. 結果の検証
	// 		if w.Code != tt.c {
	// 			t.Errorf("status code got %d, want %d", w.Code, tt.c)
	// 		}

	// 		// 必要であればレスポンスボディ(w.Body.String())の検証もここで行えます
	// 	})
	// }
}
