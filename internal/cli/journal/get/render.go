package get

import (
	"os"
	"sort"
	"strconv"
	"time"

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
	// tw.SetAutoWrapText(true)
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

	var prevDate time.Time
	for _, v := range journals {
		if v.User != userId {
			continue
		}

		date := ""
		if !prevDate.Equal(v.Date.Time) {
			date = v.Date.Format("02.01.2006")
		}

		tw.Append([]string{
			date,
			strconv.Itoa(v.Ticket),
			v.TicketTitle,
			v.Comment,
		})
	}

	return tw
}
