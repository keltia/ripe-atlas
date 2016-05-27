/*
This package is just a collection of test cases
*/
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"ripe-atlas"
	"strconv"
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
					Action: func(c *cli.Context) error {
						q, err := atlas.GetProbes()
						if err != nil {
							fmt.Printf("err: %v", err)
							os.Exit(1)
						}
						fmt.Printf("q: %#v\n", q)

						return nil
					},
				},
				{
					Name:        "info",
					Usage:       "info for one probe",
					Description: "gives info for one probe",
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "id",
							Value: 0,
							Usage: "id of the probe",
						},
					},
					Action: func(c *cli.Context) error {
						args := c.Args()
						id, _ := strconv.ParseInt(args[0], 10, 32)

						p, err := atlas.GetProbe(int(id))
						if err != nil {
							fmt.Printf("err: %v", err)
							os.Exit(1)
						}
						fmt.Printf("p: %#v\n", p)

						return nil
					},
				},
			},
		},
		{
			Name:        "ip",
			Usage:       "returns current ip",
			Description: "shorthand for getting current ip",
			Action: func(c *cli.Context) error {
				args := c.Args()
				id, _ := strconv.ParseInt(args[0], 10, 32)

				p, err := atlas.GetProbe(int(id))
				if err != nil {
					fmt.Printf("err: %v", err)
					os.Exit(1)
				}
				fmt.Printf("IPv4: %v - IPv6: %v", p.AddressV4, p.AddressV6)
				return nil
			},
		},

	}
	app.Run(os.Args)

}
