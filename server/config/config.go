package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

type AppEnv string

const (
	ENV_PRODUCTION  AppEnv = "PROD"
	ENV_DEVELOPMENT AppEnv = "DEV"
)

type Config struct {
	App     AppConfig
	Context ContextConfig
	DB      DBConfig
	Cors    CorsConfig
	Token   TokenConfig
}

func (c Config) IsProduction() bool {
	return c.App.Environment == ENV_PRODUCTION
}

type AppConfig struct {
	Version     string
	Port        int
	Environment AppEnv
	DNS         string
}

type ContextConfig struct {
	Timeout int
}

type DBConfig struct {
	Username        string
	Password        string
	Host            string
	Port            int
	Name            string
	MinConns        int // PG Pool minimum connections
	MaxConns        int // PG Pool maximum connections
	MaxConnLifetime int // PG Pool maximun connection lifetime, In Minute
}

type CorsConfig struct {
	AllowedOrigins string
	Credentials    bool
}

// Convert AllowedOrigins into Array of string
func (cc CorsConfig) ParseOrigins() ([]string, error) {
	origins := strings.Split(cc.AllowedOrigins, ",")
	// Validate Origins URL
	for _, origin := range origins {
		_, err := url.ParseRequestURI(origin)
		if err != nil {
			return nil, err
		}
	}
	return origins, nil
}

type TokenConfig struct {
	SecretKey        string
	Expiry           int    // In Hour
	AccessCookieName string // Access Token Cookie Name
}

// Load and Parse Config
func LoadConfig(path string) (*Config, error) {
	var c Config
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(path) // Respectfully from the root directory

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fail load config: %v", err.Error())
	}

	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("fail parse config: %v", err.Error())
	}

	if c.App.Environment != ENV_PRODUCTION && c.App.Environment != ENV_DEVELOPMENT {
		return nil, errors.New("Invalid app.Environtment value, expecting DEV or PROD")
	}

	return &c, nil
}
