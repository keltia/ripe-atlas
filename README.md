# ripe-atlas

* RIPE Atlas v2 API access in Go. *

[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/keltia/ripe-atlas/atlas) [![license](https://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/keltia/ripe-atlas/master/LICENSE) [![build](https://img.shields.io/travis/keltia/ripe-atlas.svg?style=flat)](https://travis-ci.org/ant0ine/ripe-atlas)

**ripe-atlas is a Go library to access the RIPE Atlas [REST API](https://atlas.ripe.net/docs/api/v2/manual/).

It features a simple CLI-based tool called `atlas` which use the library.

**Work in progress, still incomplete**

## Table of content
 
- [Features](#features)
- [Install](#install)

  Like many Go-based tools, installation is very easy
  
    go install github.com/keltia/ripe-atlas
  
  The library is fetched, compiled and installed in whichever directory is specified by `$GOPATH`.  The `atlas` binary will also be installed. 

- [API usage](#api-usage)
  - [Basics](#basics)
      - [Authentication](#auth)
	  - [Probes](#probes)
	  - [Measurements](#measurements)
  - [Applications](#applications)
  
  The `atlas` command is a command-line client for the Go API:
  
  ```
  NAME:
     atlas - RIPE Atlas cli interface
  
  USAGE:
     atlas [global options] command [command options] [arguments...]
  
  AUTHOR(S):
     Ollivier Robert <roberto@keltia.net>
  
  COMMANDS:
       ip				returns current ip
       measurements, measures, m	measurements-related keywords
       ping			ping selected address
       probes, p, pb		probe-related keywords
  
  GLOBAL OPTIONS:
     --format value, -f value	specify output format
     -v				verbose mode
     --fields value, -F value	specify which fields are wanted
     --opt-fields value, -O value	specify which optional fields are wanted
     --sort value, -S value	sort results (default: "id")
     --help, -h			show help
  ```
  
  In addition to the main `probes` and `measurements` commands, it features fast-access to common tasks like `ping`and `traceroute`.
- [External Documentation](#external-documentation)
  [Main RIPE Atlas site](https://atlas.ripe.net/)
  [REST API Documentation](https://atlas.ripe.net/docs/api/v2/manual/)
  [REST API Reference](https://atlas.ripe.net/docs/api/v2/reference/)