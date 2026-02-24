package main

import (
	"log"
	"time"

	"react-ts/backend/config"
	"react-ts/backend/controller"
	_ "react-ts/backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title	react-ts backend API
//	@version	1.0
//	@BasePath	/v1

func main() {

	// 設定を読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// ginのDefault Routerを生成
	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(corsHandler(cfg))

	// 各エンドポイントのルーティング
	c := controller.NewController()

	v1 := r.Group("/v1")
	{
		surveyors := v1.Group("/surveyors")
		{
			surveyors.GET("query", c.QuerySurveyors)
		}
	}

	// Swagger UIのルーティング
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// サーバー起動
	r.Run(":" + cfg.Port)
}

// CORSミドルウェアの設定を行いHandlerFuncを返す
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
