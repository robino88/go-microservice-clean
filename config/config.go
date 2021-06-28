package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"net/url"
	"time"
)

//Config Config struct contain specific struct to organize the configuration instead of having it
// All in one struct
type Config struct {
	Server        ServerConfig
	Commercetools CommercetoolsConfig
	Debug         bool `env:"DEBUG,required"`
}

//ServerConfig Specific config struct for all the Server configuration
type ServerConfig struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type CommercetoolsConfig struct {
	Project      string   `env:"CTP_PROJECT_KEY,required"`
	ClientId     string   `env:"CTP_CLIENT_ID,required"`
	ClientSecret string   `env:"CTP_CLIENT_SECRET,required"`
	OauthUrl     string   `env:"CTP_AUTH_URL,required"`
	ApiUrl       *url.URL `env:"CTP_API_URL,required"`
	Scopes       []string `env:"CTP_SCOPES,required"`
}

func AppConfig() *Config {
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode %s", err)
	}

	return &c
}
