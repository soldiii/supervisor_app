package server

import "github.com/soldiii/supervisor_app/internal/store"

type ServerConfig struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.StoreConfig
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr:     ":8080",
		LogLevel: "debug",
		Store:    store.NewStoreConfig(),
	}
}
