package main

import (
	"fmt"
	"github.com/jgrigorian/actions-run-times/cmd/list"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	// CLI
	app := &cli.App{
		Name:  "actions-run-times",
		Usage: "Tool to get the average run times of GitHub Actions Workflows",
		OnUsageError: func(ctx *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(ctx.App.Writer, "WRONG: %#v\n", err)
			return nil
		},
		Commands: []*cli.Command{
			///////////////////////////////////////////////////////////
			// LIST
			///////////////////////////////////////////////////////////
			{
				Name:  "list",
				Usage: "Option for listing workflows",
				Subcommands: []*cli.Command{
					{
						Name:  "workflows",
						Usage: "list workflows",
						Action: func(c *cli.Context) error {
							list.Workflows(c)
							return nil
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "owner",
								Aliases: []string{"o"},
								Usage:   "Owner of the repository",
							},
							&cli.StringFlag{
								Name:    "repo",
								Aliases: []string{"r"},
								Usage:   "Name of the repository",
							},
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
