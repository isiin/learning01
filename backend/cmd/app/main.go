package main

import (
	"log"

	"react-ts/backend/config"
	"react-ts/backend/internal/api"
	"react-ts/backend/internal/bootstrap"
)

func main() {

	// 設定を読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// 依存関係の設定
	cp := bootstrap.NewComponents()

	// サーバー起動
	api.Run(cfg, cp)
}
