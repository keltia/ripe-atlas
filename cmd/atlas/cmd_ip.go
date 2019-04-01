package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/urfave/cli"
)

var (
	// For --bare/-B
	fBare bool
)

// init injects our "ip" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ip",
		Usage:       "returns current ip",
		Description: "shorthand for getting current ip",
		Action:      cmdIP,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "B, bare",
				Usage:       "Minimal output template",
				Destination: &fBare,
			},
		},
	})

}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {

	var (
		probeID int
	)

	args := c.Args()
	if len(args) == 1 {
		probeID, _ = strconv.Atoi(args[0])
	}

	if probeID == 0 {
		if cnf.DefaultProbe == 0 {
			log.Fatal("Error: you must specify a probe ID!")
		} else {
			probeID = cnf.DefaultProbe
		}
	}

	p, err := client.GetProbe(probeID)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	// really minimalistic output, only v4
	if fBare {
		fmt.Printf("%s %s\n", p.AddressV4, p.AddressV6)
	} else {
		fmt.Printf("IPv4: %s IPv6: %s\n", p.AddressV4, p.AddressV6)
	}
	return nil
}
