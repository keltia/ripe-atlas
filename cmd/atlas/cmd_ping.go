package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/keltia/ripe-atlas"
	"log"
	"os"
)

// init injects our "ip" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ping",
		Usage:       "ping selected address",
		Description: "send echo/reply to an IP",
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
		Action: cmdPing,
	})
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdPing(c *cli.Context) error {
	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}
	args := c.Args()
	if args[0] == "" {
		log.Fatalf("Error: you must specify a hostname/IP")
	}

	addr := args[0]

	def := atlas.Definition{
		Target: addr,
	}
	defs := []atlas.Definition{}
	defs = append(defs, def)

	req := atlas.MeasurementRequest{
		Definitions: defs,
	}
	// Default set of probes
	probes := atlas.ProbeSet{
		{
			Requested: 10,
			Type: "area",
			Value: "WW",
		},
	}

	req.Probes = probes
	_, err := atlas.Ping(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	//str := res.Result.Display()
	fmt.Println()
	return nil
}

