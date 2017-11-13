# ripe-atlas

* RIPE Atlas v2 API access in Go. *

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/keltia/ripe-atlas) [![license](https://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/keltia/ripe-atlas/master/LICENSE) [![build](https://img.shields.io/travis/keltia/ripe-atlas.svg?style=flat)](https://travis-ci.org/keltia/ripe-atlas) [![Go Report Card](https://goreportcard.com/badge/github.com/keltia/ripe-atlas)](https://goreportcard.com/report/github.com/keltia/ripe-atlas)

`ripe-atlas` is a [Go](https://golang.org/) library to access the RIPE Atlas [REST API](https://atlas.ripe.net/docs/api/v2/manual/).

It features a simple CLI-based tool called `atlas` which serve both as a collection of use-cases for the library and an easy way to use it.

**Work in progress, still incomplete**

## Table of content
 
- [Features](#features)
- [Installation](#installation)
- [API usage](#api-usage)
  - [Basics](#basics)
- [CLI Application](#cli-application)
- [TODO](#todo)
- [External Documentation](#external-documentation)

## Features

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
- sslcert/tls
- traceroute

## Installation

Like many Go-based tools, installation is very easy
  
    go get github.com/keltia/ripe-atlas/cmd/...

or
  
    git clone https://github.com/keltia/ripe-atlas
    make install
    
The library is fetched, compiled and installed in whichever directory is specified by `$GOPATH`.  The `atlas` binary will also be installed (on windows, this will be called `atlas.exe`). 

You can install the dependencies with `go get`
  
- `github.com/sendgrid/rest`
- `github.com/urfave/cli`
- `github.com/naoina/toml`

To run the tests, you will also need:

- `github.com/stretchr/assert`

## API usage

You must foremost instanciate a new API client with

    client, err := atlas.NewClient(config)

where `config` is an `atlas.Config{}` struct with various options.

All API calls after that will use `client`:

    probe, err := client.GetProbe(12345)

### Basics

- Authentication
- Probes
- Measurements
- Applications

## CLI utility

The `atlas` command is a command-line client for the Go API.

### Configuration

The `atlas` utility uses a configuration file in the [TOML](https://github.com/naoina/toml) file format.

On UNIX, it is located in `$HOME/.config/ripe-atlas/config.toml` and in `%LOCALAPPDATA%\RIPE-ATLAS` on Windows. 

There are only a few parameters for now, the most important one being your API Key for autheicate against the RIPE API endpoint.

    API_key = "UUID"
    pool_size = 10
    default_probe = "YOUR-PROBE-ID"

### Usage

```
NAME:
   atlas - RIPE Atlas CLI interface

USAGE:
   atlas [global options] command [command options] [arguments...]

VERSION:
   0.11

AUTHOR:
   Ollivier Robert <roberto@keltia.net>

COMMANDS:
     credits, c                 credits-related keywords
     dns, dig, drill            send dns queries
     http, https                connect to host/IP through HTTP
     ip                         returns current ip
     keys, k, key               key-related keywords
     measurements, measures, m  measurements-related keywords
     ntp                        get time from ntp server
     ping                       ping selected address
     probes, p, pb              probe-related keywords
     results, r, res            results for one measurement
     sslcert, tlscert, tls      get TLS certificate from host/IP
     traceroute, trace          traceroute to given host/IP
     help, h                    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value, -f value      specify output format
   --debug, -D                   debug mode
   --verbose, -v                 verbose mode
   --fields value, -F value      specify which fields are wanted
   --include value, -I value     specify whether objects should be expanded
   --mine, -M                    limit output to my objects
   --opt-fields value, -O value  specify which optional fields are wanted
   --page-size value, -P value   page size for results
   --sort value, -S value        sort results
   -6, --ipv6                    Only IPv6
   -4, --ipv4                    Only IPv4
   --help, -h                    show help
   --version, -V
```
  
In addition to the main `probes` and `measurements` commands, it features fast-access to common tasks like `ping`and `traceroute`.

When looking at measurement results, it is very easy to use something like [jq](https://stedolan.github.io/jq) to properly display JSON data:

    atlas results <ID> | jq .

### TODO

- more tests (and better ones!)
- better display of results
- refactoring to reduce code duplication: done
- even more tests

### External Documentation

  - [Main RIPE Atlas site](https://atlas.ripe.net/)
  - [REST API Documentation](https://atlas.ripe.net/docs/api/v2/manual/)
  - [REST API Reference](https://atlas.ripe.net/docs/api/v2/reference/)
  