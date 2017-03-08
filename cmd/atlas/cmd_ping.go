package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
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
		Action:      cmdPing,
	})
}

// shortcuts

func preparePing(target string) (req *atlas.MeasurementRequest) {
	opts := map[string]string{
		"Type":        "ping",
		"Description": fmt.Sprintf("Ping - %s", target),
		"Target":      target,
	}

	req = atlas.NewMeasurement()
	if mycnf.WantAF == WantBoth {

		opts["AF"] = "4"
		req.AddDefinition(opts)

		opts["AF"] = "6"
		req.AddDefinition(opts)
	} else {
		opts["AF"] = mycnf.WantAF
		req.AddDefinition(opts)
	}

	// Check global parameters
	opts = checkGlobalFlags(opts)

	if fVerbose {
		displayOptions(opts)
	}

	return
}

// cmdIP is a short for displaying the IPs for one probe
func cmdPing(c *cli.Context) error {
	args := c.Args()
	if args == nil || len(args) != 1 {
		log.Fatal("Error: you must specify a hostname/IP")
	}

	addr := args[0]

	req := preparePing(addr)
	log.Printf("req=%#v", req)
	m, err := atlas.Ping(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	//str := res.Result.Display()
	fmt.Printf("m: %v\n", m)
	return nil
}
