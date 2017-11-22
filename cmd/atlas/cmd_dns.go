package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

const (
	defQueryType  = "ANY"
	defQueryClass = "IN"
	defProtocol   = "UDP"
)

var (
	eDNS0 = false
)

// checkQueryParam checks against possible list of parameters
func checkQueryParam(arg string, list map[string]bool) bool {
	_, ok := list[strings.ToUpper(arg)]
	return ok
}

// init injects our "dns" related commands/options.
func init() {
	// Fill-in the various commands
	cliCommands = append(cliCommands, cli.Command{
		Name:        "dns",
		Usage:       "send dns queries",
		Description: "send DNS queries about an host/IP/domain\n   use: <Q> [<TYPE> [<CLASS>]]",
		Aliases: []string{
			"dig",
			"drill",
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "E, edns0",
				Usage:       "use EDNS0",
				Destination: &eDNS0,
			},
			cli.BoolFlag{
				Name:        "D, disable-dnssec",
				Usage:       "Do not try to validate DNSSEC RR",
				Destination: &fDisableDNSSEC,
			},
			cli.BoolFlag{
				Name:        "C, disable-dnssec-checks",
				Usage:       "Do not try to validate DNSSEC Check by probes",
				Destination: &fBitCD,
			},
			cli.StringFlag{
				Name:        "p, protocol",
				Usage:       "Select UDP or TCP",
				Destination: &fProtocol,
			},
		},
		Action: cmdDNS,
	})
}

func prepareDNS(proto, qa, qc, qt string, do, cd bool) (req *atlas.MeasurementRequest) {
	opts := map[string]string{
		"Type":          "dns",
		"Description":   fmt.Sprintf("DNS - %s", qa),
		"Protocol":      proto,
		"Target":        qa,
		"QueryArgument": qa,
		"QueryClass":    qc,
		"QueryType":     qt,
		"SetDOBit":      boolToString(do),
		"SetCDBit":      boolToString(cd),
	}

	if eDNS0 {
		opts["UDPPayloadSize"] = "4096"
		opts["Protocol"] = "UDP"
	} else {
		opts["UDPPayloadSize"] = "512"
	}

	// Check global parameters
	opts = checkGlobalFlags(opts)

	req = client.NewMeasurement()

	if wantAF == WantBoth {

		opts["AF"] = "4"
		req.AddDefinition(opts)

		opts["AF"] = "6"
		req.AddDefinition(opts)
	} else {
		opts["AF"] = wantAF
		req.AddDefinition(opts)
	}

	if fVerbose {
		displayOptions(opts)
	}

	return
}

func cmdDNS(c *cli.Context) error {
	var (
		bitDO  = true
		bitCD  = false
		qtype  = defQueryType
		qclass = defQueryClass
		proto  = defProtocol

		addr string
	)

	args := c.Args()
	if args == nil || len(args) == 0 {
		log.Fatal("Error: you must specify at least a name")
	}

	if len(args) == 1 {
		addr = args[0]
	} else if len(args) == 2 {
		addr = args[0]
		qtype = strings.ToUpper(args[1])
	} else if len(args) == 3 {
		addr = args[0]
		qtype = strings.ToUpper(args[1])
		qclass = strings.ToUpper(args[2])
	}

	if fProtocol != "" {
		log.Printf("Use %s", fProtocol)
		proto = fProtocol
	}

	if fDisableDNSSEC {
		bitDO = false
	}

	if fBitCD {
		bitCD = true
	}

	req := prepareDNS(proto, addr, qclass, qtype, bitDO, bitCD)

	if fDebug {
		log.Printf("req=%#v", req)
	}
	m, err := client.DNS(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	//str := res.Result.Display()
	displayMeasurementID(*m)

	return nil
}
