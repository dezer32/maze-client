package authorize

import (
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/container"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "authorize",
		Short: "Get auth token",
		Run: func(*cobra.Command, []string) {
			c := container.Client()
			c.Authorize()
			container.Config().SetToken(c.Authorization)
		},
	}
}
