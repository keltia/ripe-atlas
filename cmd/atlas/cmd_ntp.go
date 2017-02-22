package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
)

// init injects our "ntp" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ntp",
		Usage:       "get time from ntp server",
		Description: "send NTP queries to an host/IP",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "6, ipv6",
				Usage:       "displays only IPv6",
				Destination: &fWant6,
			},
			cli.BoolFlag{
				Name:        "4, ipv4",
				Usage:       "displays only IPv4",
				Destination: &fWant4,
			},
		},
		Action: cmdNTP,
	})
}

func cmdNTP(c *cli.Context) error {
	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}

	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	target := args[0]
	var defs []atlas.Definition

	if fWant4 {
		def := atlas.Definition{
			AF:          4,
			Description: fmt.Sprintf("NTP v4 - %s", target),
			Type:        "ntp",
			Target:      target,
		}
		defs = append(defs, def)
	}

	if fWant6 {
		def := atlas.Definition{
			AF:          6,
			Description: fmt.Sprintf("NTP v6 - %s", target),
			Type:        "ntp",
			Target:      target,
		}
		defs = append(defs, def)
	}

	req := atlas.MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
	}

	// Default set of probes
	probes := atlas.ProbeSet{
		{
			Requested: mycnf.PoolSize,
			Type:      "area",
			Value:     "WW",
			Tags:      nil,
		},
	}

	req.Probes = probes
	log.Printf("req=%#v", req)
	//str := res.Result.Display()

	tls, err := atlas.NTP(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("NTP: %#v", tls)

	return nil
}
