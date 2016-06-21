// cmd_measures.go

package main

import (
	"github.com/codegangsta/cli"
)

// init injects our probe-related commands
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name: "measurements",
		Aliases: []string{
			"measures",
			"m",
		},
		Usage:       "measurements-related keywords",
		Description: "All the commands for measurements",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Aliases:     []string{"ls"},
				Usage:       "lists all measurements",
				Description: "displays all measurements",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "country,c",
						Usage:       "filter on country",
						Value:       "fr",
						Destination: &fCountry,
					},
					cli.StringFlag{
						Name:        "asn",
						Usage:       "filter on asn",
						Value:       "",
						Destination: &fAsn,
					},
					cli.BoolFlag{
						Name:        "A",
						Usage:       "all measurements even inactive ones",
						Destination: &fAllMeasurements,
					},
					cli.BoolFlag{
						Name:        "is-anchor",
						Usage:       "select anchor measurements",
						Destination: &fWantAnchor,
					},
				},
				Action: measurementList,
			},
			{
				Name:        "info",
				Usage:       "info for one measurement",
				Description: "gives info for one measurement",
				Action:      measurementInfo,
			},
		},
	})
}

func measurementList(c *cli.Context) error {
	return nil
}

func measurementInfo(c *cli.Context) error {
	return nil
}

func measurementCreate(c *cli.Context) error {
	return nil
}
