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
	// By default we want both
	if !want4 && !want6 {
		want6, want4 = true, true
	}
	args := c.Args()
	id, _ := strconv.ParseInt(args[0], 10, 32)

	p, err := atlas.GetProbe(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	var str string = ""

	if want4 {
		str = fmt.Sprintf("%sIPv4: %s ", str, p.AddressV4)
	}

	if want6 {
		str = fmt.Sprintf("%sIPv6: %s ", str, p.AddressV6)
	}

	fmt.Println(str)
	return nil
}
