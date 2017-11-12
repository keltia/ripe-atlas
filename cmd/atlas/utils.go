package main

import (
	"github.com/urfave/cli"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// ByAlphabet is for sorting
type ByAlphabet []cli.Command

func (a ByAlphabet) Len() int           { return len(a) }
func (a ByAlphabet) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAlphabet) Less(i, j int) bool { return a[i].Name < a[j].Name }

// checkGlobalFlags is the place to check global parameters
func checkGlobalFlags(o map[string]string) (opts map[string]string) {
	opts = o

	for k, v := range map[string]string{
		"fields":          fFieldList,
		"format":          fFormat,
		"include":         fInclude,
		"optional_fields": fOptFields,
		"page":            fPageNum,
		"page_size":       fPageSize,
		"sort":            fSortOrder,
	} {
		if v != "" {
			client.SetOption(k, v)
		}
	}

	if fFormat != "" && validateFormat(fFormat) {
		client.SetOption("format", fFormat)
	}
	return opts
}

// validateFormat allows only supported formats
func validateFormat(fmt string) bool {
	f := strings.ToLower(fmt)
	if f == "json" || f == "xml" || f == "api" || f == "txt" || f == "jsonp" {
		return true
	}
	return false
}

func displayOptions(opts map[string]string) {
	log.Println("Options:")
	for key, val := range opts {
		log.Printf("  %s: %s", key, val)
	}
}

func boolToString(k bool) string {
	if k {
		return "true"
	} else {
		return "false"
	}
}

// analyzeTarget breaks up an url into its components
func analyzeTarget(target string) (proto, site, path string, port int) {
	uri, err := url.Parse(target)
	if err != nil {
		proto = ""
		site = ""
		path = ""
		port = 0
	} else {
		proto = uri.Scheme
		if proto == "https" {
			port = 443
		}

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
