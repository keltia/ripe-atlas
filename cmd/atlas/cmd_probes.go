// cmd_probes.go

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/keltia/ripe-atlas"
	"os"
	"strconv"
)

// probeList displays all probes
func probesList(c *cli.Context) error {
	q, err := atlas.GetProbes()
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("q: %#v\n", q)

	return nil
}

// probeInfo is information about one probe
func probeInfo(c *cli.Context) error {
	args := c.Args()
	id, _ := strconv.ParseInt(args[0], 10, 32)

	p, err := atlas.GetProbe(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("p: %#v\n", p)

	return nil
}

// shortcuts

// cmdIP is a short for displaying the IPs for one probe
func cmdIP(c *cli.Context) error {
	args := c.Args()
	id, _ := strconv.ParseInt(args[0], 10, 32)

	p, err := atlas.GetProbe(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("IPv4: %v - IPv6: %v", p.AddressV4, p.AddressV6)
	return nil
}
