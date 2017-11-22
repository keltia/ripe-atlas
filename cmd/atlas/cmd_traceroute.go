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

// prepareTraceroute build the request with our parameters
func prepareTraceroute(target, protocol string, maxhops, size int) (req *atlas.MeasurementRequest) {
	opts := map[string]string{
		"Type":        "traceroute",
		"Description": fmt.Sprintf("Traceroute - %s", target),
		"Target":      target,
		"MaxHops":     fmt.Sprintf("%d", maxhops),
		"Size":        fmt.Sprintf("%d", size),
		"Protocol":    protocol,
	}

	// Try to configure -4/-6 depending on the argument to DTRT
	prepareFamily(target)

	req = client.NewMeasurement()
	if wantAF == WantBoth {

		opts["AF"] = "4"
		req.AddDefinition(opts)

		opts["AF"] = "6"
		req.AddDefinition(opts)
	} else {
		opts["AF"] = wantAF
		req.AddDefinition(opts)
	}

	return
}

func cmdTraceroute(c *cli.Context) error {
	var (
		maxHops    int
		packetSize int
		proto      string
	)

	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	target := args[0]

	proto = defProtocol
	maxHops = defMaxHops

	if fPacketSize != 0 {
		packetSize = fPacketSize
	}
	if fMaxHops != 0 {
		maxHops = fMaxHops
	}

	req := prepareTraceroute(target, proto, maxHops, packetSize)

	if fDebug {
		log.Printf("req=%#v", req)
	}
	//str := res.Result.Display()

	trc, err := client.Traceroute(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	displayMeasurementID(*trc)

	return nil
}
