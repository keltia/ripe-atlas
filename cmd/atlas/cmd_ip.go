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
		Action: cmdIP,
	})
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {
	var probeID string

	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}
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

	id, _ := strconv.ParseInt(probeID, 10, 32)

	p, err := atlas.GetProbe(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	var str string

	if fWant4 {
		str = fmt.Sprintf("%sIPv4: %s ", str, p.AddressV4)
	}

	if fWant6 {
		str = fmt.Sprintf("%sIPv6: %s ", str, p.AddressV6)
	}

	fmt.Println(str)
	return nil
}
