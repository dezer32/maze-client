package get

import (
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"

	"github.com/dezer32/maze-client/internal/client"
)

type journalSort []*client.Journal

func (v journalSort) Len() int           { return len(v) }
func (v journalSort) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v journalSort) Less(i, j int) bool { return v[j].Date.Before(v[i].Date.Time) }

func Render(journals []*client.Journal, userId int) *tablewriter.Table {
	sort.Sort(journalSort(journals))

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetRowLine(true)
	tw.SetAutoMergeCells(true)

	tw.SetHeader([]string{"date", "ticket", "title", "comment"})
	tw.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor},
	)

	tw.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
	)

	for _, v := range journals {
		if v.User != userId {
			continue
		}

		tw.Append([]string{
			v.Date.Format("02.01"),
			strconv.Itoa(v.Ticket),
			v.TicketTitle,
			v.Comment,
		})
	}

	return tw
}
