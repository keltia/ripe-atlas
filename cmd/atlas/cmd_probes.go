// cmd_probes.go

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"ripe-atlas"
	"strconv"
	"log"
)

// displayProbe display short or verbose data about a probe
func displayProbe(p *atlas.Probe, verbose bool) (res string) {
	if verbose {
		res = fmt.Sprintf("%v\n", p)
	} else {
		res = fmt.Sprintf("ID: %d Country: %s IPv4: %s IPv6: %s Descr: %s\n",
			p.ID,
			p.CountryCode,
			p.AddressV4,
			p.AddressV6,
			p.Description)
	}
	return
}

func displayAllProbes(pl *[]atlas.Probe, verbose bool) (res string) {
	res = ""
	for _, p := range *pl {
		// Do we want the inactive probes as well?
		if p.AddressV4 == "" && p.AddressV6 == "" {
			if !fAllProbes {
				continue
			}
		}
		res += displayProbe(&p, verbose)
	}
	return
}

// probeList displays all probes
func probesList(c *cli.Context) error {
	opts := make(map[string]string)

	if fCountry != "" {
		opts["country_code"] = fCountry
	}

	if fAsn != "" {
		opts["asn"] = fAsn
	}

	q, err := atlas.GetProbes(opts)
	if err != nil {
		log.Printf("GetProbes err: %v - q:%v", err, q)
		os.Exit(1)
	}
	log.Printf("Got %d probes with %v\n", len(q), opts)
	fmt.Print(displayAllProbes(&q, fVerbose))

	return nil
}

// probeInfo is information about one probe
func probeInfo(c *cli.Context) error {
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
	fmt.Print(displayProbe(p, fVerbose))

	return nil
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
