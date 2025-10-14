package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
)

type Environment struct {
	AppEnv string `env:"APP_ENV"`
}

type Jwt struct {
	AccessSecret        string `env:"JWT_ACCESS_SECRET" env-default:"secret"`
	RefreshSecret       string `env:"JWT_REFRESH_SECRET" env-default:"secret"`
	AccessExpireMinutes string `env:"JWT_ACCESS_EXPIRE_MINUTES" env-default:"15"`
	RefreshExpireHours  string `env:"JWT_REFRESH_EXPIRE_HOURS" env-default:"168"`
	PrivatePemPath      string `env:"JWT_PRIVATE_PEM_PATH"`
	PublicPemPath       string `env:"JWT_PUBLIC_PEM_PATH"`
}

type Http struct {
	FullHttpHost string `env:"FULL_HTTP_HOST"`
	HttpPort     string `env:"HTTP_PORT"`
}

type Config struct {
	Environment Environment `env:"environment"`
	Jwt         Jwt         `env:"jwt"`
	Http        Http        `env:"http"`
}

func Make() *Config {
	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading .env file %v", err)
	}

	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Panicf("Error reading env file %v", err)
	}
	return &config
}

func (s *Config) JwtAccessSecret() string {
	return s.Jwt.AccessSecret
}

func (s *Config) JwtRefreshSecret() string {
	return s.Jwt.RefreshSecret
}

func (s *Config) JwtAccessExpireMinutes() string {
	return s.Jwt.AccessExpireMinutes
}

func (s *Config) JwtRefreshExpireHours() string {
	return s.Jwt.RefreshExpireHours
}

func (s *Config) JwtPrivatePemPath() string {
	return s.Jwt.PrivatePemPath
}

func (s *Config) JwtPublicPemPath() string {
	return s.Jwt.PublicPemPath
}

func (s *Config) FullHttpHost() string {
	return s.Http.FullHttpHost
}

func (s *Config) HttpPort() string {
	return s.Http.HttpPort
}
