package main

import "github.com/codegangsta/cli"

// init injects our "sslcert" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "sslcert",
		Aliases:     []string{"tlscert", "tls"},
		Usage:       "get TLS certificate from host/IP",
		Description: "connect and fetch TLS certificate from host/IP",
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
		Action: cmdTLSCert,
	})
}

func cmdTLSCert(c *cli.Context) error {
	return nil
}
