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
	OAuth    OAuthConfig    `mapstructure:"oauth"`
}

type OAuthConfig struct {
	ClientID     string          `mapstructure:"client_id"`
	ClientSecret string          `mapstructure:"client_secret"`
	RedirectURL  string          `mapstructure:"redirect_url"`
	Scopes       []string        `mapstructure:"scopes"`
	Endpoint     oauth2.Endpoint `mapstructure:"endpoint"`
	UserinfoURL  string          `mapstructure:"userinfo_url"`
}

func (o *OAuthConfig) ToOAuth2Config() oauth2.Config {
	return oauth2.Config{
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		RedirectURL:  o.RedirectURL,
		Scopes:       o.Scopes,
		Endpoint:     o.Endpoint,
	}
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
