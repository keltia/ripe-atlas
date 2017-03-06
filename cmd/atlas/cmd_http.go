package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
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
				Value:       "GET",
				Destination: &fHTTPMethod,
			},
			cli.StringFlag{
				Name:        "U, user-agent",
				Usage:       "Override User-Agent.",
				Destination: &fUserAgent,
			},
			cli.StringFlag{
				Name:        "V, version",
				Usage:       "Set a specific HTTP version.",
				Destination: &fHTTPVersion,
			},
		},
		Action: cmdHTTP,
	})
}

func prepareHTTP(target, path string, port int) (req *atlas.MeasurementRequest) {
	opts := map[string]string{
		"Type":        "http",
		"Method":      fHTTPMethod,
		"UserAgent":   fUserAgent,
		"Version":     fHTTPVersion,
		"Description": fmt.Sprintf("HTTP - %s", target),
		"Target":      target,
		"Path":        path,
		"Port":        fmt.Sprintf("%d", port),
	}

	req = atlas.NewMeasurement()
	if mycnf.WantAF == WantBoth {

		opts["AF"] = "4"
		req.AddDefinition(opts)

		opts["AF"] = "6"
		req.AddDefinition(opts)
	} else {
		opts["AF"] = mycnf.WantAF
		req.AddDefinition(opts)
	}

	return
}

func cmdHTTP(c *cli.Context) error {
	args := c.Args()
	if len(args) == 0 {
		log.Fatal("Error: you must specify a hostname/site!")
	}

	// We expect target to be using [http|https]://<site>[:port]/path
	target := args[0]

	proto, site, path, port := analyzeTarget(target)
	if proto == "" || site == "" {
		return fmt.Errorf("Invalid URL: %s", target)
	}

	log.Printf("Target: %s://%s:%d%s", proto, site, port, path)

	req := prepareHTTP(site, path, port)

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
