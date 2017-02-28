package main

import (
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"fmt"
	"log"
	"os"
)

// init injects our "http" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "http",
		Aliases:     []string{"https"},
		Usage:       "connect to host/IP through HTTP",
		Description: "send HTTP queries to an host/IP",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "X, method",
				Usage:       "Use this method instead of default GET",
				Destination: &fHTTPMethod,
			},
			cli.StringFlag{
				Name:        "U, user-agent",
				Usage:       "Override User-Agent.",
				Destination: &fUserAgent,
			},
			cli.StringFlag{
				Name:        "V, version",
				Usage:       "Set a specific version.",
				Destination: &fVersion,
			},
			cli.BoolFlag{
				Name:        "6, ipv6",
				Usage:       "displays only IPv6",
				Destination: &fWant6,
			},
			cli.BoolFlag{
				Name:        "4, ipv4",
				Usage:       "displays only IPv4",
				Destination: &fWant4,
			},
		},
		Action: cmdHTTP,
	})
}

func cmdHTTP(c *cli.Context) error {
	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}

	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	// We expect target to be using [http|https]://<site>[:port]/path
	target := args[0]
	var defs []atlas.Definition

	_, site, path, port := analyzeTarget(target)

	if fWant4 {
		def := atlas.Definition{
			AF:          4,
			Description: fmt.Sprintf("HTTP v4 - %s", target),
			Type:        "http",
			Method:      "GET",
			Target:      site,
			Port:        port,
			Path:        path,
		}
		if fHTTPMethod != "" {
			def.Method = fHTTPMethod
		}
		if fUserAgent != "" {
			def.UserAgent = fUserAgent
		}

		defs = append(defs, def)
	}

	if fWant6 {
		def := atlas.Definition{
			AF:          6,
			Description: fmt.Sprintf("HTTP v6 - %s", target),
			Type:        "http",
			Method:      "GET",
			Target:      site,
			Port:        port,
			Path:        path,
		}
		if fHTTPMethod != "" {
			def.Method = fHTTPMethod
		}
		if fUserAgent != "" {
			def.UserAgent = fUserAgent
		}
		if fVersion != "" {
			def.Version = fVersion
		} else {
			def.Version = atlasVersion
		}

		defs = append(defs, def)
	}

	req := atlas.MeasurementRequest{
		Definitions: defs,
		IsOneoff:    true,
	}
	// Default set of probes
	probes := atlas.ProbeSet{
		{
			Requested: mycnf.PoolSize,
			Type:      "area",
			Value:     "WW",
			Tags:      nil,
		},
	}

	req.Probes = probes
	log.Printf("req=%#v", req)
	//str := res.Result.Display()

	http, err := atlas.HTTP(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("HTTP: %#v", http)
	return nil

}
