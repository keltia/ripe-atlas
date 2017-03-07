package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
)

// init injects our "ip" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ip",
		Usage:       "returns current ip",
		Description: "shorthand for getting current ip",
		Action:      cmdIP,
	})
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {
	var probeID string

	args := c.Args()
	if len(args) == 0 {
		if mycnf.DefaultProbe == 0 {
			log.Fatal("Error: you must specify a probe ID!")
		} else {
			probeID = fmt.Sprintf("%d", mycnf.DefaultProbe)
		}
	} else {
		probeID = args[0]
	}

	id, _ := strconv.Atoi(probeID)

	p, err := atlas.GetProbe(id)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	fmt.Printf("IPv4: %s IPv6: %s\n", p.AddressV4, p.AddressV6)
	return nil
}
