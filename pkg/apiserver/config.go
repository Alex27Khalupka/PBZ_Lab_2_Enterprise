package apiserver

import "github.com/Alex27Khalupka/PBZ_Lab_2_Enterprise/pkg/store"

type Config struct{
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

func NewConfig() *Config{
	return &Config{
		BindAddr: ":8080",
		Store:	  store.NewConfig(),
	}
}

