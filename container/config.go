package container

import (
	"github.com/dezer32/maze-client/internal/core/config"
)

var cfg *config.Config

func Config() *config.Config {
	if cfg == nil {
		cfg = config.NewConfig(".env")
	}

	return cfg
}
