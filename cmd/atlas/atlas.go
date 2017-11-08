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
	fWantMine = true

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

	fDebug      bool
	fVerbose    bool
	fWantAnchor bool

	fMaxHops    int
	fPacketSize int

	mycnf *Config

	cliCommands []cli.Command

	client *atlas.Client
)

const (
	atlasVersion = "0.11"
	MyName       = "ripe-atlas"

	// WantBoth is the way to ask for both IPv4 & IPv6.
	WantBoth = "64"

	// Want4 only 4
	Want4 = "4"
	// Want6 only 6
	Want6 = "6"
)

// -4 & -6 are special, if neither is specified, then we turn both as true
// Check a few other things while we are here
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

	// Logical
	if fDebug {
		fVerbose = true
		log.Printf("config: %#v", mycnf)
	}

	// Various messages
	if fVerbose {
		if mycnf.APIKey != "" {
			log.Printf("Found API key!")
		} else {
			log.Printf("No API key!")
		}

		if mycnf.DefaultProbe != 0 {
			log.Printf("Found default probe: %d\n", mycnf.DefaultProbe)
		}
	}

	return nil
}

// main is the starting point (and everything)
func main() {
	cli.VersionFlag = cli.BoolFlag{Name: "version, V"}

	cli.VersionPrinter = func(c *cli.Context) {
		log.Printf("API wrapper: %s Atlas API: %s\n", c.App.Version, atlas.GetVersion())
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
			Name:        "debug,D",
			Usage:       "debug mode",
			Destination: &fDebug,
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

	// Load main configuration
	mycnf, err = LoadConfig("")
	if err != nil {
		if fVerbose {
			log.Printf("No configuration file found.")
		}
	}

	// Check whether we have proxy authentication (from a separate config file)
	auth, err := setupProxyAuth()
	if err != nil {
		if fVerbose {
			log.Printf("Invalid or no proxy auth credentials")
		}
	}

	// Wondering whether to move to the Functional options pattern
	// cf. https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
	client, err = atlas.NewClient(atlas.Config{
		APIKey:       mycnf.APIKey,
		DefaultProbe: mycnf.DefaultProbe,
		PoolSize:     mycnf.PoolSize,
		ProxyAuth:    auth,
		Verbose:      fVerbose,
	})

	// No need to continue if this fails
	if err != nil {
		log.Fatalf("Error creating the Atlas client: %v", err)
	}

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
