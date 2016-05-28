/*
This package is just a collection of test cases
*/
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

// main is the starting point (and everything)
func main() {
	app := cli.NewApp()
	app.Name = "atlas"
	app.Usage = "RIPE Atlas cli interface"
	app.Author = "Ollivier Robert <roberto@keltia.net>"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name: "probes",
			Aliases: []string{
				"p",
				"pb",
			},
			Usage:       "probe-related keywords",
			Description: "All the commands for probes",
			Subcommands: []cli.Command{
				{
					Name:        "list",
					Aliases:     []string{"ls"},
					Usage:       "lists all probes",
					Description: "displays all probes",
					Action:      probesList,
				},
				{
					Name:        "info",
					Usage:       "info for one probe",
					Description: "gives info for one probe",
					Action:      probeInfo,
				},
			},
		},
		{
			Name:        "ip",
			Usage:       "returns current ip",
			Description: "shorthand for getting current ip",
			Action:      cmdIP,
		},
	}
	app.Run(os.Args)

}
