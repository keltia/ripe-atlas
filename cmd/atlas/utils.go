package main

import (
	"fmt"
	"github.com/keltia/ripe-atlas"
	"github.com/urfave/cli"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// ByAlphabet is for sorting
type ByAlphabet []cli.Command

func (a ByAlphabet) Len() int           { return len(a) }
func (a ByAlphabet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlphabet) Less(i, j int) bool { return a[i].Name < a[j].Name }

func boolToString(k bool) string {
	if k {
		return "true"
	}
	return "false"
}

// analyzeTarget breaks up an url into its components
func analyzeTarget(target string) (proto, site, path string, port int) {
	uri, err := url.Parse(target)
	if err != nil {
		// This is no URL, make one
		proto = "https"
		site = target
		path = "/"
		port = 443
	} else {
		proto = uri.Scheme

		// might be host:port
		sp := strings.Split(uri.Host, ":")
		if len(sp) == 2 {
			port, _ = strconv.Atoi(sp[1])
			site = sp[0]
		} else {
			site = uri.Host
		}

		path = uri.Path
		// Path can't be null
		if path == "" {
			path = "/"
		}
	}
	return
}

const (
	// Hostname is a domain or machine name
	hostname = iota
	ipv4
	ipv6
)

// checkArgumentType checks whether we have a domain/host or an IPv4/v6 address
func checkArgumentType(arg string) int {
	ip := net.ParseIP(arg)
	if ip != nil {
		if ip.To4() == nil {
			return ipv6
		}
		return ipv4
	}
	return hostname
}

// prepareFamily sets Want4 and/or Want6 depending on the argument
func prepareFamily(arg string) {
	switch checkArgumentType(arg) {
	case ipv4:
		wantAF = Want4
	case ipv6:
		wantAF = Want6
	default:
		wantAF = WantBoth
	}
}

// displayMeasurementID display result of measurement requests.
func displayMeasurementID(list atlas.MeasurementResp) {
	fmt.Println("Measurements created:")
	for _, m := range list.Measurements {
		fmt.Printf("%d\n", m)
	}
	fmt.Printf(`
Use the following command to retrieve results in JSON:

  atlas r <id>

`)
}

// debug displays only if fDebug is set
func debug(str string, a ...interface{}) {
	if fDebug {
		log.Printf(str, a...)
	}
}

// verbose displays only if fVerbose is set
func verbose(str string, a ...interface{}) {
	if fVerbose {
		log.Printf(str, a...)
	}
}
