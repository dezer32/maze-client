package journal

import (
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/internal/cli/journal/get"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "journal",
		Short: "Journal instruments.",
	}

	cmd.AddCommand(
		get.NewCommand(),
	)

	return cmd
}
