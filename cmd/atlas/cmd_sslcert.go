package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
)

// init injects our "sslcert" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name: "sslcert",
		Aliases: []string{
			"tlscert",
			"tls",
		},
		Usage:       "get TLS certificate from host/IP",
		Description: "connect and fetch TLS certificate from host/IP",
		Action:      cmdTLSCert,
	})
}

func cmdTLSCert(c *cli.Context) (err error) {
	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}

	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	target := args[0]
	var defs []atlas.Definition

	if fWant4 {
		def := atlas.Definition{
			AF:          4,
			Description: fmt.Sprintf("TLS v4 - %s", target),
			Type:        "sslcert",
			Target:      target,
		}
		defs = append(defs, def)
	}

	if fWant6 {
		def := atlas.Definition{
			AF:          6,
			Description: fmt.Sprintf("TLS v6 - %s", target),
			Type:        "sslcert",
			Target:      target,
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

	tls, err := atlas.SSLCert(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("TLS: %#v", tls)
	return
}
