package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database databaseConfig
	Server   serverConfig
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
	URL string
}

func (c *databaseConfig) validate() error {
	if c.URL == "" {
		return errors.New("database url is required")
	}
	return nil
}

type serverConfig struct {
	Host string
	Port int
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

	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.url", "")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			log.Fatal("config file not found")
		}
	}

	return &Config{
		Database: databaseConfig{
			URL: viper.GetString("database.url"),
		},
		Server: serverConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
	}
}
