package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"ripe-atlas"
	"strconv"
)

// init injects our "ip" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "ip",
		Usage:       "returns current ip",
		Description: "shorthand for getting current ip",
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
		Action: cmdIP,
	})
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {
	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}
	args := c.Args()
	if args[0] == "" {
		log.Fatalf("Error: you must specify a probe ID!")
	}

	id, _ := strconv.ParseInt(args[0], 10, 32)

	p, err := atlas.GetProbe(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	var str string = ""

	if fWant4 {
		str = fmt.Sprintf("%sIPv4: %s ", str, p.AddressV4)
	}

	if fWant6 {
		str = fmt.Sprintf("%sIPv6: %s ", str, p.AddressV6)
	}

	fmt.Println(str)
	return nil
}
