package config

import (
	"errors"
	"log"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	Database databaseConfig `mapstructure:"database"`
	Server   serverConfig   `mapstructure:"server"`
	OAuth    *oauth2.Config `mapstructure:"oauth"`
}

func (c *Config) Validate() error {
	if err := c.Database.validate(); err != nil {
		return err
	}
	if err := c.Server.validate(); err != nil {
		return err
	}
	return nil
}

type databaseConfig struct {
	URL string `mapstructure:"url"`
}

func (c *databaseConfig) validate() error {
	if c.URL == "" {
		return errors.New("database url is required")
	}
	return nil
}

type serverConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (c *serverConfig) validate() error {
	if c.Host == "" {
		return errors.New("server host is required")
	}
	if c.Port == 0 {
		return errors.New("server port is required")
	}
	return nil
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")

	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.url", "")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			log.Fatal("config file not found")
		}
	}

	var cfg Config
	err := viper.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.MatchName = func(mapKey, fieldName string) bool {
			normalize := func(s string) string {
				return strings.ToLower(strings.ReplaceAll(s, "_", ""))
			}
			return normalize(mapKey) == normalize(fieldName)
		}
	})
	if err != nil {
		log.Fatalf("unable to decode config: %v", err)
	}

	return &cfg
}
