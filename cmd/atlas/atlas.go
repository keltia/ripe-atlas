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
	fWant4 bool
	fWant6 bool

	// True by default
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

	WantBoth = "64"
	Want4    = "4"
	Want6    = "6"
)

// -4 & -6 are special, if neither is specified, then we turn both as true
func finalcheck(c *cli.Context) error {
	if fWant4 {
		mycnf.WantAF = Want4
	}

	if fWant6 {
		mycnf.WantAF = Want6
	}

	// Both are fine
	if fWant4 && fWant6 {
		mycnf.WantAF = WantBoth
	}

	// So is neither — common case
	if !fWant4 && !fWant6 {
		mycnf.WantAF = WantBoth
	}

	return nil
}

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

	// Ensure -4 & -6 are treated properly
	app.Before = finalcheck

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

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
