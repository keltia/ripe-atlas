// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

// +build windows !unix

package main

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	basedir = filepath.Join(os.Getenv("%LOCALAPPDATA%"),
		"RIPE-ATLAS",
	)
)

// Check the parameter for either tag or filename
func checkName(file string) (str string) {
	// Just use default location
	if file == "" {
		str = filepath.Join(basedir, "config.toml")
	} else {
		// Full path, begin by \\ or N: MUST have .toml
		bfile := []byte(file)
		if bfile[1] == ':' || strings.HasPrefix(file, "\\\\") {
			if !strings.HasSuffix(file, ".toml") {
				str = ""
			} else {
				str = file
			}
			return
		}

		// If ending with .toml, take it literally
		if strings.HasSuffix(file, ".toml") {
			str = file
			return
		}
	}
	return
}
