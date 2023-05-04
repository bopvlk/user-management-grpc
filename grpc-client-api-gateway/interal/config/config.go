package config

import (
	"git.foxminded.com.ua/grpc/grpc-client-api-gateway/interal/apperrors"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	GRPCPort   string `env:"GRPC_PORT"`
	HTTPPort   string `env:"HTTP_PORT"`
	SigningKey string `env:"SIGNING_KEY"`
	TokenTtl   int    `env:"TOKEN_TTL"`
}

func InitConfig() (config *Config, err error) {
	config = &Config{}
	if err := env.Parse(config); err != nil {
		return nil, apperrors.EnvConfigParseError.AppendMessage(err)
	}
	return
}
