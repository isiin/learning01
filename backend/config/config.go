package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config はアプリケーションの設定を保持します。
type Config struct {
	AllowOrigin string
	Port        string
}

// Load は .env ファイルと環境変数から設定を読み込みます。
func Load() (Config, error) {
	var cfg Config

	// .envファイルが存在しなくてもエラーにはしません。
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return cfg, fmt.Errorf("failed to load .env file: %w", err)
		}
	}

	cfg.AllowOrigin = os.Getenv("ALLOW_ORIGIN")
	if cfg.AllowOrigin == "" {
		cfg.AllowOrigin = "http://localhost:5173"
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg, nil
}
