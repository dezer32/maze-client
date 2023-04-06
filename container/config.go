package container

import (
	"github.com/dezer32/maze-client/internal/core/config"
)

var cfg *config.Config

func ReloadConfig(cfgFile string) *config.Config {
	cfg = config.NewConfig(cfgFile)

	return cfg
}

func Config() *config.Config {
	if cfg == nil {
		cfg = config.NewConfig(".env")
	}

	return cfg
}
