package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"time"
)

//Config Config struct contain specific struct to organize the configuration instead of having it
// All in one struct
type Config struct {
	Server ServerConfig
}

//ServerConfig Specifc config struct for all the Server configuration
type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode %s", err)
	}

	return &c
}
