package api

import (
	"react-ts/backend/config"
	_ "react-ts/backend/docs"
	"react-ts/backend/internal/api/middleware"
	v1 "react-ts/backend/internal/api/v1"
	"react-ts/backend/internal/bootstrap"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run はサーバーを設定し起動します
func Run(cfg config.Config, cp *bootstrap.Components) {
	// ginのDefault Routerを生成
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(middleware.CorsHandler(cfg))

	// 各エンドポイントのルーティング
	v1.Route(r, cp)

	// Swagger UIのルーティング
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// サーバー起動
	r.Run(":" + cfg.Port)
}
