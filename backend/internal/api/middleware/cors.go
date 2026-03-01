package middleware

import (
	"react-ts/backend/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsHandler はCORSミドルウェアの設定を行いHandlerFuncを返します
func CorsHandler(cfg config.Config) gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			cfg.AllowOrigin,
		},
		// 許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
		},
		// 許可したいHTTPヘッダー
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})
}
