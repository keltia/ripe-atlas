package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
)

// init injects our "ntp" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ntp",
		Usage:       "get time from ntp server",
		Description: "send NTP queries to an host/IP",
		Action:      cmdNTP,
	})
}

// prepareTraceroute build the request with our parameters
func prepareNTP(target string) (req *atlas.MeasurementRequest) {
	opts := map[string]string{
		"Type":        "ntp",
		"Description": fmt.Sprintf("NTP - %s", target),
		"Target":      target,
	}

	// Check global parameters
	opts = checkGlobalFlags(opts)

	if fVerbose {
		displayOptions(opts)
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

func cmdNTP(c *cli.Context) error {
	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	target := args[0]

	req := prepareNTP(target)
	debug("req=%#v", req)
	//str := res.Result.Display()

	ntp, err := client.NTP(req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	displayMeasurementID(*ntp)

	return nil
}
