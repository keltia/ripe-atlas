// cmd_probes.go

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"ripe-atlas"
	"strconv"
)

// init injects our probe-related commands
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name: "probes",
		Aliases: []string{
			"p",
			"pb",
		},
		Usage:       "probe-related keywords",
		Description: "All the commands for probes",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Aliases:     []string{"ls"},
				Usage:       "lists all probes",
				Description: "displays all probes",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "country,c",
						Usage:       "filter on country",
						Value:       "fr",
						Destination: &fCountry,
					},
					cli.StringFlag{
						Name:        "asn",
						Usage:       "filter on asn",
						Value:       "",
						Destination: &fAsn,
					},
					cli.BoolFlag{
						Name:        "A",
						Usage:       "all probes even inactive ones",
						Destination: &fAllProbes,
					},
					cli.BoolFlag{
						Name:        "is-anchor",
						Usage:       "select anchor probes",
						Destination: &fWantAnchor,
					},
				},
				Action: probesList,
			},
			{
				Name:        "info",
				Usage:       "info for one probe",
				Description: "gives info for one probe",
				Action: probeInfo,
			},
		},
	})
}

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

	if fWantAnchor {
		opts["is_anchor"] = "true"
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

