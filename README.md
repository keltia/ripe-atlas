# ripe-atlas

* RIPE Atlas v2 API access in Go. *

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/keltia/ripe-atlas) [![license](https://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/keltia/ripe-atlas/master/LICENSE) [![build](https://img.shields.io/travis/keltia/ripe-atlas.svg?style=flat)](https://travis-ci.org/keltia/ripe-atlas)

**ripe-atlas is a Go library to access the RIPE Atlas [REST API](https://atlas.ripe.net/docs/api/v2/manual/).

It features a simple CLI-based tool called `atlas` which use the library.

**Work in progress, still incomplete**

## Table of content
 
- [Features](#features)
- [Install](#install)
- [API usage](#api-usage)
  - [Basics](#basics)
      - [Authentication](#auth)
	  - [Probes](#probes)
	  - [Measurements](#measurements)
  - [Applications](#applications)
- [External Documentation](#external-documentation)

## [#features]: Features

I am trying to implement the full REST API in Go.  The API itself is not particularly complex but the settings and parameters are.

The following topic are available:

- probes

  you can query one probe or ask for a list of probes with a few criterias
  
- measurements

  you can create and list measurements.
  
- results

  every measurement has a URI in the result json that points to the actual results. This fetch and display them. 

In addition to these major commands, there are a few shortcut commands (see below):

- dns
- http
- ip
- keys
- ntp
- ping
- sslcert
- traceroute

## Installation

  Like many Go-based tools, installation is very easy
  
    go get github.com/keltia/ripe-atlas

  or
  
    git clone https://github.com/keltia/ripe-atlas
    go install ./cmd/...
    
  The library is fetched, compiled and installed in whichever directory is specified by `$GOPATH`.  The `atlas` binary will also be installed. 

  Dependencies:
## API usage

### Configuration

This package uses a configuration file in the [TOML](https://github.com/naoina/toml) file format located by default in `$HOME/.ripe-atlas/config.toml`.

There are only a few parameters for now, the most important one being your API Key for autheicate against the RIPE API endpoint.

    API_key = "UUID"
    pool_size = 10
    default_probe = "YOUR-PROBE-ID"

### Basics

- Authentication
- Probes
- Measurements
- Applications

  The `atlas` command is a command-line client for the Go API:
  
  ```
  NAME:
     atlas - RIPE Atlas cli interface
  
  USAGE:
     atlas [global options] command [command options] [arguments...]
  
  AUTHOR(S):
     Ollivier Robert <roberto@keltia.net>
  
  COMMANDS:
     dns			send dns queries
     keys           key management
     http, https		connect to host/IP through HTTP
     ip				returns current ip
     measurements, measures, m	measurements-related keywords
     ntp			get time from ntp server
     ping			ping selected address
     probes, p, pb		probe-related keywords
     sslcert, tlscert, tls	get TLS certificate from host/IP
     traceroute, trace		traceroute to given host/IP

  GLOBAL OPTIONS:
     --format value, -f value	specify output format
     --verbose, -v		verbose mode
     --fields value, -F value	specify which fields are wanted
     --opt-fields value, -O value	specify which optional fields are wanted
     --sort value, -S value	sort results (default: "id")
     --help, -h			show help
  ```
  
  In addition to the main `probes` and `measurements` commands, it features fast-access to common tasks like `ping`and `traceroute`.

### TODO

- more tests
- better display of results
- implementation behind many keywords
- even more tests

### External Documentation

  - [Main RIPE Atlas site](https://atlas.ripe.net/)
  - [REST API Documentation](https://atlas.ripe.net/docs/api/v2/manual/)
  - [REST API Reference](https://atlas.ripe.net/docs/api/v2/reference/)
  