package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/dezer32/maze-client/internal/cli"
)

func main() { os.Exit(run()) }

func run() int {
	logrus.Info("Running...")

	cmd := cli.NewCommand()
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("Can't run cmd.")
		return 1
	}

	logrus.Info("Canceled.")
	return 0
}
