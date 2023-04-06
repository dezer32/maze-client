package cli

import (
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/internal/cli/authorize"
	"github.com/dezer32/maze-client/internal/cli/journal"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "maze-cli",
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
