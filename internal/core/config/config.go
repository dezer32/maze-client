package config

import (
	"github.com/spf13/viper"

	"github.com/dezer32/maze-client/internal/core/logger"
)

type Authorization struct {
	Token  string `mapstructure:"AUTHORIZATION_TOKEN" json:"token"`
	UserId int    `mapstructure:"AUTHORIZATION_USER" json:"user"`
}

type Config struct {
	Url           string `mapstructure:"URL"`
	Username      string `mapstructure:"USERNAME"`
	Password      string `mapstructure:"PASSWORD"`
	Authorization *Authorization
	viper         *viper.Viper
}

func NewConfig(cfg string) *Config {
	v := viper.New()
	v.SetConfigType("env")
	v.SetConfigFile(cfg)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		logger.Log.WithError(err).Fatal("Can't read config file.")
	}

	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		logger.Log.WithError(err).Fatal("Can't unmarshal config.")
	}

	config.Authorization = &Authorization{}
	if err := v.Unmarshal(config.Authorization); err != nil {
		logger.Log.WithError(err).Fatal("Can't unmarshal auth config.")
	}

	config.viper = v

	return config
}

func (c *Config) SetToken(authorization *Authorization) {
	c.Authorization = authorization
	c.viper.Set("AUTHORIZATION_TOKEN", c.Authorization.Token)
	c.viper.Set("AUTHORIZATION_USER", c.Authorization.UserId)
	err := c.viper.WriteConfig()
	if err != nil {
		logger.Log.WithError(err).Error("Can't save config.")
	}
}
