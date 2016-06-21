// cmd_measures.go

package main

import (
	"github.com/codegangsta/cli"
	"log"
	"strconv"
	"github.com/keltia/ripe-atlas"
	"fmt"
	"os"
)

// init injects our probe-related commands
func init() {
	cliCommands = append(cliCommands, cli.Command{
		Name: "measurements",
		Aliases: []string{
			"measures",
			"m",
		},
		Usage:       "measurements-related keywords",
		Description: "All the commands for measurements",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Aliases:     []string{"ls"},
				Usage:       "lists all measurements",
				Description: "displays all measurements",
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
						Usage:       "all measurements even inactive ones",
						Destination: &fAllMeasurements,
					},
					cli.BoolFlag{
						Name:        "is-anchor",
						Usage:       "select anchor measurements",
						Destination: &fWantAnchor,
					},
				},
				Action: measurementsList,
			},
			{
				Name:        "info",
				Usage:       "info for one measurement",
				Description: "gives info for one measurement",
				Action:      measurementInfo,
			},
		},
	})
}

func displayMeasurement(m *atlas.Measurement, verbose bool) (res string) {
	if verbose {
		res = fmt.Sprintf("%v\n", m)
	} else {
		res = fmt.Sprintf("ID: %d type: %s description: %s\n", m.ID, m.Type, m.Description)
	}
	return
}

func displayAllMeasurements(ml *[]atlas.Measurement, verbose bool) (res string) {
	res = ""
	for _, m := range *ml {
		res += displayMeasurement(&m, verbose)
	}
	return
}


func measurementsList(c *cli.Context) error {
	opts := make(map[string]string)

	if fCountry != "" {
		opts["country_code"] = fCountry
	}

	if fAsn != "" {
		opts["asn"] = fAsn
	}

	// Check global parameters
	opts = checkGlobalFlags(opts)

	q, err := atlas.GetMeasurements(opts)
	if err != nil {
		log.Printf("GetMeasurements err: %v - q:%v", err, q)
		os.Exit(1)
	}
	log.Printf("Got %d measurements with %v\n", len(q), opts)
	fmt.Print(displayAllMeasurements(&q, fVerbose))

	return nil
	return nil
}

func measurementInfo(c *cli.Context) error {
	args := c.Args()
	if args[0] == "" {
		log.Fatalf("Error: you must specify a measurement ID!")
	}

	id, _ := strconv.ParseInt(args[0], 10, 32)

	p, err := atlas.GetMeasurement(int(id))
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Print(displayMeasurement(p, fVerbose))

	return nil
}

func measurementCreate(c *cli.Context) error {
	return nil
}
