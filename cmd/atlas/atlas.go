/*
This package is just a collection of use-case for the various aspects of the RIPE API.
Consider this both as an example on how to use the API and a testing tool for the API wrapper.
*/
package main

import (
	"fmt"
	"github.com/keltia/proxy"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

var (
	// CLI specific options
	fDebug   bool
	fLogfile string
	fVerbose bool

	// See flag.go for details

	// Global API options
	fFieldList string
	fFormat    string
	fInclude   string
	fOptFields string
	fPageNum   string
	fPageSize  string
	fSortOrder string
	fWantMine  bool

	// Probe-specific ones
	fAllProbes bool
	fIsAnchor  bool
	fTags      string

	// Common measurement ones
	fAllMeasurements bool
	fAsn             string
	fAsnV4           string
	fAsnV6           string
	fCountry         string
	fProtocol        string
	fMeasureType     string
	fWant4           bool
	fWant6           bool

	// Create measurements
	fBillTo    string
	fIsOneOff  bool
	fStartTime string
	fStopTime  string

	// HTTP
	fHTTPMethod  string
	fUserAgent   string
	fHTTPVersion string

	// DNS
	fBitCD         bool
	fDisableDNSSEC bool

	// Traceroute
	fMaxHops    int
	fPacketSize int

	// ProbeSet parameters
	fPoolSize  int
	fAreaType  string
	fAreaValue string

	// Our configuration file
	cnf *Config

	// All possible commands
	cliCommands []cli.Command

	client *atlas.Client

	// Our tiple-valued synthesis of fWant4/fWant6
	wantAF string
)

const (
	atlasVersion = "0.31"
	// MyName is the application name
	MyName = "ripe-atlas"

	// WantBoth is the way to ask for both IPv4 & IPv6.
	WantBoth = "64"

	// Want4 only 4
	Want4 = "4"
	// Want6 only 6
	Want6 = "6"
)

func openlog(fn string) *log.Logger {
	fh, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error: can not open logfile %s: %v", fn, err)
	}

	mylog := log.New(fh, "", log.LstdFlags)
	verbose("Logfile: %s %#v", fn, mylog)

	return mylog
}

// -4 & -6 are special, if neither is specified, then we turn both as true
// Check a few other things while we are here
func finalcheck(c *cli.Context) error {

	var (
		err   error
		mylog *log.Logger
	)

	// Load main configuration
	cnf, err = LoadConfig("")
	if err != nil {
		verbose("No configuration file found.")
	}

	// Logical
	if fDebug {
		fVerbose = true
		verbose("config: %#v", cnf)
	}

	// Various messages
	if cnf.APIKey != "" {
		verbose("Found API key!")
	} else {
		verbose("No API key!")
	}

	if cnf.DefaultProbe != 0 {
		verbose("Found default probe: %d\n", cnf.DefaultProbe)
	}

	// Allow overwrite of a few parameters
	if fPoolSize != 0 {
		cnf.ProbeSet.PoolSize = fPoolSize
	}
	if fAreaType != "" {
		cnf.ProbeSet.Type = fAreaType
	}
	if fAreaValue != "" {
		cnf.ProbeSet.Value = fAreaValue
	}

	// Check whether we have proxy authentication (from a separate config file)
	auth, err := proxy.SetupProxyAuth()
	if err != nil {
		verbose("Invalid or no proxy auth credentials")
	}

	// If we want a logfile, open one for the API to log into
	if fLogfile != "" {
		mylog = openlog(fLogfile)
	}

	// Check whether the -C <country> was specified then override configuration
	if fCountry != "" {
		cnf.ProbeSet.Type = "country"
		cnf.ProbeSet.Value = fCountry
	}

	// Check whether tags are specified
	if fTags != "" {
		cnf.ProbeSet.Tags = fTags
	}

	// Wondering whether to move to the Functional options pattern
	// cf. https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
	client, err = atlas.NewClient(atlas.Config{
		APIKey:       cnf.APIKey,
		DefaultProbe: cnf.DefaultProbe,
		IsOneOff:     fIsOneOff,
		PoolSize:     cnf.ProbeSet.PoolSize,
		AreaType:     cnf.ProbeSet.Type,
		AreaValue:    cnf.ProbeSet.Value,
		Tags:         cnf.ProbeSet.Tags,
		ProxyAuth:    auth,
		Verbose:      fVerbose,
		Log:          mylog,
	})

	// No need to continue if this fails
	if err != nil {
		log.Fatalf("Error creating the Atlas client: %v", err)
	}

	if fWant4 {
		wantAF = Want4
	}

	if fWant6 {
		wantAF = Want6
	}

	// Both are fine
	if fWant4 && fWant6 {
		wantAF = WantBoth
	}

	// So is neither — common case
	if !fWant4 && !fWant6 {
		wantAF = WantBoth
	}

	debug("wantAF=%s", wantAF)
	return nil
}

// main is the starting point (and everything)
func main() {
	cli.VersionFlag = cli.BoolFlag{Name: "version, V"}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("API wrapper: %s Atlas API: %s\n", c.App.Version, atlas.GetVersion())
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
			Usage:       "specify output format (NOT IMPLEMENTED)",
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
			Name:        "include,I",
			Usage:       "specify whether objects should be expanded",
			Destination: &fInclude,
		},
		cli.StringFlag{
			Name:        "logfile,L",
			Usage:       "specify a log file",
			Destination: &fLogfile,
		},
		cli.BoolFlag{
			Name:        "mine,M",
			Usage:       "limit output to my objects",
			Destination: &fWantMine,
		},
		cli.StringFlag{
			Name:        "opt-fields,O",
			Usage:       "specify which optional fields are wanted",
			Destination: &fOptFields,
		},
		cli.StringFlag{
			Name:        "page-size,P",
			Usage:       "page size for results",
			Destination: &fPageSize,
		},

		cli.StringFlag{
			Name:        "sort,S",
			Usage:       "sort results",
			Destination: &fSortOrder,
		},
		cli.BoolTFlag{
			Name:        "1,is-oneoff",
			Usage:       "one-time measurement",
			Destination: &fIsOneOff,
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
		// These are not global parameters but it makes sense to define them only once
		// and not in every cmd_* files.
		cli.IntFlag{
			Name:        "pool-size,N",
			Usage:       "Number of probes to request",
			Destination: &fPoolSize,
		},
		cli.StringFlag{
			Name:        "area-type",
			Usage:       "Set type for probes (area, country, etc.)",
			Destination: &fAreaType,
		},
		cli.StringFlag{
			Name:        "area-value",
			Usage:       "Value for the probe set (WW, West, etc.)",
			Destination: &fAreaValue,
		},
		cli.StringFlag{
			Name:        "country,C",
			Usage:       "Short cut to specify a country",
			Destination: &fCountry,
		},
		cli.StringFlag{
			Name:        "tags,T",
			Usage:       "Include/exclude tags for probesets",
			Destination: &fTags,
		},
	}

	// Ensure -4 & -6 are treated properly & initialization is done
	app.Before = finalcheck

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
