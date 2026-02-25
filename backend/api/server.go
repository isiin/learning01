package api

import (
	"time"

	v1 "react-ts/backend/api/v1"
	"react-ts/backend/config"
	_ "react-ts/backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run はサーバーを設定し起動します
func Run(cfg config.Config) {
	// ginのDefault Routerを生成
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(corsHandler(cfg))

	// 各エンドポイントのルーティング
	v1.Route(r)

	// Swagger UIのルーティング
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// サーバー起動
	r.Run(":" + cfg.Port)
}

// corsHandler はCORSミドルウェアの設定を行いHandlerFuncを返します
func corsHandler(cfg config.Config) gin.HandlerFunc {
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
