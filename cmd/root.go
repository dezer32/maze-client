package cmd

import (
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/cmd/authorize"
	"github.com/dezer32/maze-client/cmd/journal"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "maze-cmd",
		Short:         "Client for maze api. FxBackOffice.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(
		authorize.NewCommand(),
		journal.NewCommand(),
	)

	return cmd
}
