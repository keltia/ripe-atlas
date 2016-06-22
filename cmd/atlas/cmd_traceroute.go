package main

import "github.com/codegangsta/cli"

// init injects our "traceroute" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "traceroute",
		Aliases:     []string{"trace"},
		Usage:       "traceroute to given host/IP",
		Description: "Send Traceroute queries to an host/IP",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "ipv6",
				Usage:       "displays only IPv6",
				Destination: &fWant6,
			},
			cli.BoolFlag{
				Name:        "ipv4",
				Usage:       "displays only IPv4",
				Destination: &fWant4,
			},
		},
		Action: cmdTraceroute,
	})
}

func cmdTraceroute(c *cli.Context) error {
	return nil
}

