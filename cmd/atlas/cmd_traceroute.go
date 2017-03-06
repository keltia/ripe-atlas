package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
)

const (
	defMaxHops = 30
)

// init injects our "traceroute" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "traceroute",
		Aliases:     []string{"trace"},
		Usage:       "traceroute to given host/IP",
		Description: "Send Traceroute queries to an host/IP",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "max_hops, m",
				Usage:       "Specify a maximum number of hops",
				Destination: &fMaxHops,
			},
			cli.IntFlag{
				Name:        "packet_size, s",
				Usage:       "Sends packets this size",
				Destination: &fPacketSize,
			},
			cli.StringFlag{
				Name:        "p, protocol",
				Usage:       "Select UDP or TCP",
				Destination: &fProtocol,
			},
		},
		Action: cmdTraceroute,
	})
}

func cmdTraceroute(c *cli.Context) error {
	var (
		maxHops    int
		packetSize int
		proto      string
	)

	// By default do both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}

	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	target := args[0]
	var defs []atlas.Definition

	proto = defProtocol
	maxHops = defMaxHops

	if fPacketSize != 0 {
		packetSize = fPacketSize
	}
	if fMaxHops != 0 {
		maxHops = fMaxHops
	}

	if fWant4 {
		def := atlas.Definition{
			AF:          4,
			Description: fmt.Sprintf("Traceroute - %s", target),
			Type:        "traceroute",
			Target:      target,
			Protocol:    proto,
			MaxHops:     maxHops,
			Size:        packetSize,
		}
		defs = append(defs, def)
	}

	if fWant6 {
		def := atlas.Definition{
			AF:          6,
			Description: fmt.Sprintf("Traceroute6 - %s", target),
			Type:        "traceroute",
			Target:      target,
			Protocol:    proto,
			MaxHops:     maxHops,
			Size:        packetSize,
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

	trc, err := atlas.Traceroute(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("NTP: %#v", trc)

	return nil
}
