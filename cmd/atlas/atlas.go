/*
This package is just a collection of test cases
*/
package main

import (
	"github.com/codegangsta/cli"
	"os"
	"sort"
)

var (
	// flags
	fWant4 bool
	fWant6 bool
	fAllProbes bool
	fAsn string
	fCountry string
	fFieldList string
	fOptFields string
	fSortOrder string
	fVerbose bool
	fWantAnchor bool

	cliCommands []cli.Command
)

type ByAlphabet []cli.Command

func (a ByAlphabet) Len() int           { return len(a) }
func (a ByAlphabet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlphabet) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Check global parameters
func checkGlobalFlags(o map[string]string) (map[string]string) {
	opts := o
	if fSortOrder != "" {
		opts["sort"] = fSortOrder
	}

	if fFieldList != "" {
		opts["fields"] = fFieldList
	}

	if fOptFields != "" {
		opts["optional_fields"] = fOptFields
	}
	return opts
}

// main is the starting point (and everything)
func main() {
	app := cli.NewApp()
	app.Name = "atlas"
	app.Usage = "RIPE Atlas cli interface"
	app.Author = "Ollivier Robert <roberto@keltia.net>"
	app.Version = "0.0.1"
	app.HideVersion = true

	// General flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "v",
			Usage: "verbose mode",
			Destination: &fVerbose,
		},
		cli.StringFlag{
			Name: "fields,F",
			Usage: "specify which fields are wanted",
			Destination: &fFieldList,
		},
		cli.StringFlag{
			Name: "opt-fields,O",
			Usage: "specify which optional fields are wanted",
			Destination: &fOptFields,
		},
		cli.StringFlag{
			Name:        "sort,S",
			Usage:       "sort results",
			Value:       "id",
			Destination: &fSortOrder,
		},
	}

	sort.Sort(ByAlphabet(cliCommands))
	app.Commands = cliCommands
	app.Run(os.Args)
}
