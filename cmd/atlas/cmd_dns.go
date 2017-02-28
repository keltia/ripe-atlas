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
	eDns0 = false
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
				Name:        "6, ipv6",
				Usage:       "displays only IPv6",
				Destination: &fWant6,
			},
			cli.BoolFlag{
				Name:        "4, ipv4",
				Usage:       "displays only IPv4",
				Destination: &fWant4,
			},
			cli.BoolFlag{
				Name:        "E, edns0",
				Usage:       "use EDNS0",
				Destination: &eDns0,
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
			/*			cli.StringFlag{
							Name:        "t, qtype",
							Usage:       "Select the query type",
							Destination: &defQueryType,
						},
						cli.StringFlag{
							Name:        "c, qclass",
							Usage:       "Select the query class",
							Destination: &defQueryClass,
						},*/
			cli.StringFlag{
				Name:        "p, protocol",
				Usage:       "Select UDP or TCP",
				Destination: &fProtocol,
			},
		},
		Action: cmdDNS,
	})
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

	// By default we want both
	if !fWant4 && !fWant6 {
		fWant6, fWant4 = true, true
	}
	args := c.Args()
	if args == nil || len(args) == 0 {
		log.Fatal("Error: you must specify at least a name")
	}

	qtype = defQueryType
	qclass = defQueryClass
	proto = defProtocol

	if len(args) == 1 {
		addr = args[0]
	} else if len(args) == 2 {
		addr = args[0]
		qtype = args[1]
	} else if len(args) == 3 {
		addr = args[0]
		qtype = args[1]
		qclass = args[2]
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

	var defs []atlas.Definition

	if fWant4 {
		def := atlas.Definition{
			AF:            4,
			Description:   fmt.Sprintf("DNS v4 - %s %s %s", addr, qtype, qclass),
			Type:          "dns",
			Protocol:      proto,
			QueryArgument: addr,
			QueryClass:    qclass,
			QueryType:     qtype,
			SetDOBit:      bitDO,
			SetCDBit:      bitCD,
		}
		if eDns0 {
			def.UDPPayloadSize = 4096
			def.Protocol = "UDP"
		}
		defs = append(defs, def)
	}

	if fWant6 {
		def := atlas.Definition{
			AF:            6,
			Description:   fmt.Sprintf("DNS v6 - %s %s %s", addr, qtype, qclass),
			Type:          "dns",
			Protocol:      proto,
			QueryArgument: addr,
			QueryClass:    qclass,
			QueryType:     qtype,
			SetDOBit:      bitDO,
			SetCDBit:      bitCD,
		}
		if eDns0 {
			def.UDPPayloadSize = 4096
			def.Protocol = "UDP"
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
	m, err := atlas.DNS(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}

	//str := res.Result.Display()
	fmt.Printf("m: %v\n", m)

	return nil
}
