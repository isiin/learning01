package main

import (
	"log"

	"react-ts/backend/api"
	"react-ts/backend/config"
)

func main() {

	// 設定を読み込み
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// サーバー起動
	api.Run(cfg)
}
