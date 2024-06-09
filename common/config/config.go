package config

import (
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME,default=tracerstudy-tracer-service"`
	Port        Port
	SIAK_API    SIAK_API
	JWT         JWTConfig
}

type Port struct {
	GRPC string `env:"PORT_GRPC,default=8081"`
	REST string `env:"PORT_REST,default=8080"`
}

type SIAK_API struct {
	URL string `env:"SIAK_API_URL"`
	KEY string `env:"SIAK_API_KEY"`
}

type JWTConfig struct {
	JwtSecretKey  string        `env:"JWT_SECRET_KEY"`
	TokenDuration time.Duration `env:"JWT_DURATION,default=30m"`
}

func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "ERROR: [NewConfig] Error while decoding env")
	}

	return &config, nil
}
