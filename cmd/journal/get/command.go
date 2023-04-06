package get

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/container"
	"github.com/dezer32/maze-client/internal/core/logger"
)

func NewCommand() *cobra.Command {
	var fromDate time.Time
	var toDate time.Time
	from := time.Now().Add(-3 * 24 * time.Hour).Format("2006-01-02")
	to := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get journal info",
		PreRun: func(*cobra.Command, []string) {
			var err error
			fromDate, err = time.Parse("2006-01-02", from)
			if err != nil {
				logger.Log.WithError(err).Fatal("Can't parse date from.")
			}

			toDate, err = time.Parse("2006-01-02", to)
			if err != nil {
				logger.Log.WithError(err).Fatal("Can't parse to date.")
			}
		},
		Run: func(*cobra.Command, []string) {
			c := container.Client()
			cfg := container.Config()

			journal := c.Journal(fromDate, toDate)
			Render(journal, cfg.Authorization.UserId).Render()

		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&from, "from", "f", from, "from date")
	flags.StringVarP(&to, "to", "t", to, "from date")

	return cmd
}
