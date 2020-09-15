package store

type Config struct{
	DatabaseURL string
}

func NewConfig() *Config{
	return &Config{}
}

