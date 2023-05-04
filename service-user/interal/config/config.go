package config

import (
	"git.foxminded.com.ua/grpc/service-user/interal/apperrors"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	GRCPPort        string `env:"GRCP_PORT"`
	DBPassword      string `env:"DB_PASSWORD"`
	DBAddr          string `env:"DB_ADDR"`
	DBName          string `env:"DB_NAME"`
	DBHost          string `env:"DB_HOST"`
	MigrateFileAddr string `env:"MIGR_FILE_ADDR"`
}

func InitConfig() (config *Config, err error) {
	config = &Config{}

	if err := env.Parse(config); err != nil {
		return nil, apperrors.ConfigReadErr.AppendMessage(err)
	}
	return
}
