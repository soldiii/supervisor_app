package store

type StoreConfig struct {
	DatabaseURL string `toml:"database_url"`
}

func NewStoreConfig() *StoreConfig {
	return &StoreConfig{}
}
