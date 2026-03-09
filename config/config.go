package config

import (
	"os"
)

type Config struct {
	MasterKey string
	AccessKey string
}

func NewConfig() *Config {
	masterKey := os.Getenv("X_Master_Key")
	if masterKey == "" {
		panic("Нет ключа шифрования")
	}

	accessKey := os.Getenv("X_Access_Key")
	if accessKey == "" {
		panic("Нет ключа шифрования")
	}
	return &Config{
		MasterKey: masterKey,
		AccessKey: accessKey,
	}
}
