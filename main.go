package main

import (
	"os"

	"github.com/dezer32/maze-client/cmd"
	"github.com/dezer32/maze-client/internal/core/logger"
)

func main() { os.Exit(run()) }

func run() int {
	logger.Log.Info("Running...")

	if err := cmd.NewCommand().Execute(); err != nil {
		logger.Log.WithError(err).Fatal("Can't run cmd.")
		return 1
	}

	logger.Log.Info("Canceled.")
	return 0
}
