/*
This package is just a collection of test cases
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

	fHTTPMethod	 string
	fUserAgent   string
	fVersion     string

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
	atlasVersion = "0.9"
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
	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
