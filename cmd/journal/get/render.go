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

func Render(journals []*client.Journal, userId int, show []string) *tablewriter.Table {
	sort.Sort(journalSort(journals))

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetRowLine(true)
	tw.SetAutoMergeCells(true)

	tw.SetHeader(show)

	headerColors, columnColors := getColors(show)

	tw.SetHeaderColor(headerColors...)
	tw.SetColumnColor(columnColors...)

	for _, v := range journals {
		if v.User != userId {
			continue
		}

		line := make([]string, len(show))
		for i, s := range show {
			switch s {
			case "date":
				line[i] = v.Date.Format("02.01")
				break
			case "ticket":
				line[i] = strconv.Itoa(v.Ticket)
				break
			case "title":
				line[i] = v.TicketTitle
				break
			case "comment":
				line[i] = v.Comment
				break
			}
		}

		tw.Append(line)
	}

	return tw
}

func getColors(show []string) ([]tablewriter.Colors, []tablewriter.Colors) {
	headerColors := make([]tablewriter.Colors, len(show))
	columnColors := make([]tablewriter.Colors, len(show))
	for i, s := range show {
		switch s {
		case "date":
			headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
			columnColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
			break
		case "ticket":
			headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor}
			columnColors[i] = tablewriter.Colors{tablewriter.FgHiRedColor}
			break
		case "title":
			headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor}
			columnColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor}
			break
		case "comment":
			headerColors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor}
			columnColors[i] = tablewriter.Colors{tablewriter.FgCyanColor}
			break
		}
	}

	return headerColors, columnColors
}
