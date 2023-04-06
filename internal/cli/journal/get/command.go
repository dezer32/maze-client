package get

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/dezer32/maze-client/container"
)

func NewCommand() *cobra.Command {
	from := time.Now().Add(-24 * time.Hour).Format("2006-01-02")
	to := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get journal info",
		Run: func(*cobra.Command, []string) {
			c := container.Client()

			fromDate, err := time.Parse("2006-01-01", from)
			if err != nil {
				logrus.WithError(err).Fatal("Can't parse date from.")
			}

			toDate, err := time.Parse("2006-01-01", to)
			if err != nil {
				logrus.WithError(err).Fatal("Can't parse to date.")
			}
			c.Journal(fromDate, toDate)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&from, "from", "f", from, "from date")
	flags.StringVarP(&to, "to", "t", to, "from date")

	return cmd
}
