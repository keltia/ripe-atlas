/*
This package is just a collection of use-case for the various aspects of the RIPE API.
Consider this both as an example on how to use the API and a testing tool for the API wrapper.
*/
package main

import (
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

var (
	// flags
	fWant4    bool
	fWant6    bool
	fWantMine bool

	fAllProbes       bool
	fAllMeasurements bool

	fAsn         string
	fCountry     string
	fFieldList   string
	fFormat      string
	fOptFields   string
	fProtocol    string
	fSortOrder   string
	fMeasureType string

	fHTTPMethod  string
	fUserAgent   string
	fHTTPVersion string

	fBitCD         bool
	fDisableDNSSEC bool

	fVerbose    bool
	fWantAnchor bool

	fMaxHops    int
	fPacketSize int

	mycnf *atlas.Config

	cliCommands []cli.Command
)

const (
	atlasVersion = "0.10"
)

// main is the starting point (and everything)
func main() {
	cli.VersionFlag = cli.BoolFlag{Name: "version, V"}

	cli.VersionPrinter = func(c *cli.Context) {
		log.Printf("API wrapper: %s Atlas CLI: %s\n", c.App.Version, atlas.GetVersion())
	}

	app := cli.NewApp()
	app.Name = "atlas"
	app.Usage = "RIPE Atlas CLI interface"
	app.Author = "Ollivier Robert <roberto@keltia.net>"
	app.Version = atlasVersion
	//app.HideVersion = true

	// General flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "format,f",
			Usage:       "specify output format",
			Destination: &fFormat,
		},
		cli.BoolFlag{
			Name:        "verbose,v",
			Usage:       "verbose mode",
			Destination: &fVerbose,
		},
		cli.StringFlag{
			Name:        "fields,F",
			Usage:       "specify which fields are wanted",
			Destination: &fFieldList,
		},
		cli.StringFlag{
			Name:        "opt-fields,O",
			Usage:       "specify which optional fields are wanted",
			Destination: &fOptFields,
		},
		cli.StringFlag{
			Name:        "sort,S",
			Usage:       "sort results",
			Destination: &fSortOrder,
		},
		cli.BoolFlag{
			Name:        "6, ipv6",
			Usage:       "Only IPv6",
			Destination: &fWant6,
		},
		cli.BoolFlag{
			Name:        "4, ipv4",
			Usage:       "Only IPv4",
			Destination: &fWant4,
		},
	}

	var err error

	mycnf, err = atlas.LoadConfig("ripe-atlas")
	if mycnf.APIKey != "" && err == nil {
		atlas.SetAuth(mycnf.APIKey)
		log.Printf("Found API key!")
	} else {
		log.Printf("No API key!")
	}
	if mycnf.DefaultProbe != 0 && err == nil {
		log.Printf("Found default probe: %d\n", mycnf.DefaultProbe)
	}

	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
