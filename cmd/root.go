package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/cmd/authorize"
	"github.com/dezer32/maze-client/cmd/journal"
	"github.com/dezer32/maze-client/container"
)

func NewCommand() *cobra.Command {
	cfg := ".env"
	cmd := &cobra.Command{
		Use:           "maze-cmd",
		Short:         "Client for maze api. FxBackOffice.",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(*cobra.Command, []string) {
			container.ReloadConfig(cfg)
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVarP(&cfg, "env", "e", cfg, "Path to .env")

	cmd.AddCommand(
		authorize.NewCommand(),
		journal.NewCommand(),
	)

	return cmd
}
