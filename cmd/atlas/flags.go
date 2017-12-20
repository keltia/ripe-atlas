package main

import (
	"log"
	"strings"
)

/*
Flags management.

There are two types of flags for commands
- global (sort, page_size)
- local (dns-related ones, ping)related ones and so forth)

Outcome is the same in all cases, a given flag is transformed into a query flag

XXX checking of options is very limited as the API server does that anyway.
*/

var (
	globalFlags = map[string]*string{
		"fields":          &fFieldList,
		"format":          &fFormat,
		"include":         &fInclude,
		"optional_fields": &fOptFields,
		"page":            &fPageNum,
		"page_size":       &fPageSize,
		"sort":            &fSortOrder,
	}

	commonFlags = map[string]*string{
		"asn":          &fAsn,
		"asn_v4":       &fAsnV4,
		"asn_v6":       &fAsnV6,
		"country_code": &fCountry,
		"type":         &fMeasureType,
	}
)

// checkGlobalFlags is the place to check global parameters
func checkGlobalFlags(o map[string]string) (opts map[string]string) {

	opts = mergeOptions(o, globalFlags)

	if fFormat != "" && validateFormat(fFormat) {
		opts["format"] = fFormat
	}

	// Boolean flag
	if fWantMine {
		opts["mine"] = "true"
	}

	return opts
}

// mergeOptions does the obvious thing
func mergeOptions(to map[string]string, from map[string]*string) map[string]string {
	o := to
	for k, v := range from {
		if *v != "" {
			o[k] = *v
		}
	}
	return o
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
