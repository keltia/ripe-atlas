// cmd_probes.go

package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
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
						Name:        "country, c",
						Usage:       "filter on country",
						Destination: &fCountry,
					},
					cli.StringFlag{
						Name:        "asn",
						Usage:       "filter on asn",
						Destination: &fAsn,
					},
					cli.BoolFlag{
						Name:        "A, all",
						Usage:       "all probes even inactive ones",
						Destination: &fAllProbes,
					},
					cli.BoolFlag{
						Name:        "is-anchor",
						Usage:       "select anchor probes",
						Destination: &fIsAnchor,
					},
				},
				Action: probesList,
			},
			{
				Name:        "info",
				Usage:       "info for one probe",
				Description: "gives info for one probe",
				Action:      probeInfo,
			},
		},
	})
}

// displayProbe display short or verbose data about a probe
func displayProbe(p *atlas.Probe, verbose bool) (res string) {
	if verbose {
		res = fmt.Sprintf("%v\n", p)
	} else {
		res = fmt.Sprintf("ID: %d Country: %s ASN4: %d IPv4: %s IPv6: %s Descr: %s\n",
			p.ID,
			p.CountryCode,
			p.AsnV4,
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

// prepareTraceroute build the request with our parameters
func prepareProbes() map[string]string {
	opts := make(map[string]string)

	opts = mergeOptions(opts, commonFlags)

	log.Printf("opts: %#v", opts)
	// Check global parameters
	opts = checkGlobalFlags(opts)

	if wantAF != WantBoth {
		opts["AF"] = wantAF
	}

	log.Printf("opts: %#v", opts)

	if fVerbose {
		displayOptions(opts)
	}

	return opts
}

// probeList displays all probes
func probesList(c *cli.Context) error {
	opts := prepareProbes()

	q, err := client.GetProbes(opts)
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
	if len(args) == 0 {
		log.Fatal("Error: you must specify a probe ID!")
	}

	id, _ := strconv.Atoi(args[0])
	p, err := client.GetProbe(id)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Print(displayProbe(p, fVerbose))

	return nil
}
