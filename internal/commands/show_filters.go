package commands

import (
	"fmt"
	"sort"

	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/pkg/report"
)

// ShowFiltersCommand configures the command name, flags, and action.
var ShowFiltersCommand = cli.Command{
	Name:  "filters",
	Usage: "Displays a search filter overview with examples",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "md, m",
			Usage: "renders valid Markdown",
		},
	},
	Action: showFiltersAction,
}

// showFiltersAction displays a search filter overview with examples.
func showFiltersAction(ctx *cli.Context) error {
	rows, cols := form.Report(&form.SearchPhotos{})

	sort.Slice(rows, func(i, j int) bool {
		if rows[i][1] == rows[j][1] {
			return rows[i][0] < rows[j][0]
		} else {
			return rows[i][1] < rows[j][1]
		}
	})

	fmt.Println(report.Table(rows, cols, ctx.Bool("md")))

	return nil
}
